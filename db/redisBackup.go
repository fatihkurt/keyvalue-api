package db

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

const BACKUP_FILE_TIME_FORMAT = "2006-01-02 15:04:05"
const BACKUP_FILE_SUFFIX = "-data.json"

func BackupInterval(interval time.Duration) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				Backup()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// Poor redis persistence replacement
func Backup() {
	fmt.Println("Backup started")
	client := RedisClient()

	ctx := context.Background()

	// res, err := client.Save(ctx).Result()

	keys, err := client.Keys(ctx, "*").Result()

	if err != nil {
		log.Fatal(err)
	}

	values, err := client.MGet(ctx, keys...).Result()

	if err != nil {
		log.Fatal(err)
	}

	var allValues []string
	for i, value := range values {
		item := []string{keys[i], fmt.Sprint(value)}
		allValues = append(allValues, strings.Join(item, ":"))
	}

	valuesJSON, err := json.Marshal(allValues)

	if err != nil {
		log.Fatal(err)
	}

	file, err := getBackupFilePath(time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}

	err = writeFile(valuesJSON, file)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bacup created successfully!")
	}
}

func Restore() {
	filepath, err := getMostRecentBackupFile()

	if err != nil {
		fmt.Println("Cannot restore backup")
		return
	}
	fmt.Println("Restore started!")

	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	var items []string
	err = json.Unmarshal(file, &items)

	if err != nil {
		log.Fatal(err)
	}

	client := RedisClient()
	ctx := context.Background()

	client.FlushDB(ctx)

	restoredRowCount := 0
	for _, item := range items {
		parts := strings.Split(item, ":")
		key := parts[0]
		val := strings.Join(parts[1:], "")
		_, err := client.Set(ctx, key, val, 0).Result()
		if err != nil {
			log.Println("Cannot set data:" + key)
		}
		restoredRowCount++
	}
	log.Printf("%d items restored from %s!", restoredRowCount, filepath)
}

func getBackupFilePath(t time.Time) (string, error) {
	tempFilePath, err := getTempFilePath()
	if err != nil {
		return "", err
	}
	filepath := tempFilePath + t.Format(BACKUP_FILE_TIME_FORMAT) + BACKUP_FILE_SUFFIX
	return filepath, nil
}

func getTempFilePath() (string, error) {
	tempPathStoreFile := "tmp/redisbackuptemppath.txt"
	file, err := ioutil.ReadFile(tempPathStoreFile) // TODO Persistence
	if err != nil {
		tempDirLoc, err := ioutil.TempDir("", "redisbackup")
		if err != nil {
			return "", err
		}
		if err = writeFile([]byte(tempDirLoc), tempPathStoreFile); err != nil {
			fmt.Printf("Failed creating file: %s", err)
			panic(err)
		}

		return tempDirLoc, nil
	}
	return string(file), nil
}

func getMostRecentBackupFile() (string, error) {
	fp, err := getTempFilePath()
	if err != nil {
		return "", err
	}
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		return "", err
	}

	var mostRecentFileTime *time.Time
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), BACKUP_FILE_SUFFIX) {
			parts := strings.Split(file.Name(), BACKUP_FILE_SUFFIX)
			fileTime, err := time.Parse(BACKUP_FILE_TIME_FORMAT, parts[0])
			if err != nil {
				fmt.Println(file.Name() + " is not a valid backup file")
				continue
			}
			if mostRecentFileTime == nil {
				mostRecentFileTime = &fileTime
			} else if mostRecentFileTime.Before(fileTime) {
				mostRecentFileTime = &fileTime
			}
		}
	}

	if mostRecentFileTime == nil {
		return "", errors.New("no backup file exists")
	}

	fp, err = getBackupFilePath(*mostRecentFileTime)
	if err != nil {
		return "", err
	}
	return fp, nil
}

func writeFile(fileContent []byte, filepath string) error {

	file, err := os.Create(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)

	_, err = datawriter.Write(fileContent)
	if err != nil {
		return err
	}
	datawriter.Flush()
	return nil
}
