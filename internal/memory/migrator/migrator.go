package migrator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SashaMelva/web-service-gin/internal/memory/connection"
	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func RunMigrationsPg(s *connection.StorageConnection, filePath string) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)

	}
	fmt.Println(executablePath)
	currentDir := filepath.Dir(executablePath)
	fmt.Println("Текущая директория исполняемого файла:", currentDir)

	if err := goose.Up(s.StorageDb, filePath); err != nil {
		return err
	}

	return nil
}
