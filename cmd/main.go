package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	userHomeDir, err := PegarUserHomeDir()
	if err != nil {
		fmt.Println("Failed to get user home directory:", err)
		return
	}

	arquivos, err := pegarArquivos(userHomeDir)
	if err != nil {
		fmt.Println("Failed to get files:", err)
		return
	}

	organizarArquivos(userHomeDir, arquivos)

}

func pegarArquivos(userHomeDir string) ([]os.DirEntry, error) {

	allArchivesDownload, err := os.ReadDir(filepath.Join(userHomeDir, "Downloads"))
	if err != nil {
		fmt.Println("Error reading Downloads directory:", err)
		return nil, err
	}

	return allArchivesDownload, nil
}

func organizarArquivos(userHomeDir string, arquivos []os.DirEntry) {
	for _, arquivo := range arquivos {

		if arquivo.IsDir() {
			continue
		}

		extensao := filepath.Ext(arquivo.Name())
		var destino string

		switch extensao {
		case ".png", ".jpg", ".jpeg", ".gif":
			destino = filepath.Join(userHomeDir, "Downloads", "Imagens")
		case ".mp4", ".mkv":
			destino = filepath.Join(userHomeDir, "Downloads", "Videos")
		case ".mp3", ".ogg":
			destino = filepath.Join(userHomeDir, "Downloads", "Musicas")
		case ".pdf":
			destino = filepath.Join(userHomeDir, "Downloads", "Documentos")
		case ".xlsx":
			destino = filepath.Join(userHomeDir, "Downloads", "Planilhas")
		case ".svg":
			destino = filepath.Join(userHomeDir, "Downloads", "ImagensVetoriais")
		default:
			destino = filepath.Join(userHomeDir, "Downloads", "Outros")
		}

		pastaDestinoErr := os.MkdirAll(destino, 0755)
		if pastaDestinoErr != nil {
			log.Fatalf("Erro ao criar a pasta de destino: %v", pastaDestinoErr)
		}
		err := os.Rename(
			filepath.Join(userHomeDir, "Downloads", arquivo.Name()),
			filepath.Join(destino, arquivo.Name()),
		)
		if err != nil {
			log.Fatalf("Erro ao mover o arquivo: %v", err)
		}
	}
}

func PegarUserHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}
