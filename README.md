# 📁 Organizador de Pastas

Um organizador automático de arquivos para sua pasta de Downloads, desenvolvido em Go.

## 📋 Descrição

Este projeto organiza automaticamente os arquivos da sua pasta de Downloads em subpastas categorizadas por tipo de arquivo (imagens, vídeos, documentos, etc.).

## ✨ Funcionalidades

- 🖼️ **Imagens**: `.png`, `.jpg`, `.jpeg`, `.gif`
- 🎬 **Vídeos**: `.mp4`, `.mkv`
- 🎵 **Músicas**: `.mp3`, `.ogg`
- 📄 **Documentos**: `.pdf`
- 📊 **Planilhas**: `.xlsx`
- 🎨 **Imagens Vetoriais**: `.svg`
- 📦 **Outros**: Arquivos não categorizados

## 🚀 Como Usar

### Pré-requisitos

- Go 1.16 ou superior instalado

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/MarkHiarley/organizador-de-pasta.git
cd organizador-de-pasta
```

2. Execute o programa:
```bash
go run cmd/main.go
```

### Compilar para diferentes sistemas

#### Linux/macOS:
```bash
go build -o organizador cmd/main.go
./organizador
```

#### Windows:
```bash
go build -o organizador.exe cmd/main.go
organizador.exe
```

#### Compilação cruzada:

**Para Windows (a partir do Linux/macOS):**
```bash
GOOS=windows GOARCH=amd64 go build -o organizador.exe cmd/main.go
```

**Para Linux (a partir do Windows/macOS):**
```bash
GOOS=linux GOARCH=amd64 go build -o organizador cmd/main.go
```

**Para macOS (a partir do Windows/Linux):**
```bash
GOOS=darwin GOARCH=amd64 go build -o organizador cmd/main.go
```

## 📂 Estrutura de Pastas Criadas

Após a execução, sua pasta Downloads terá a seguinte estrutura:

```
Downloads/
├── Imagens/
├── Videos/
├── Musicas/
├── Documentos/
├── Planilhas/
├── ImagensVetoriais/
└── Outros/
```

## 🔧 Como Funciona

1. **Detecção do Home Directory**: O programa detecta automaticamente o diretório home do usuário (funciona em Windows, Linux e macOS)
2. **Leitura dos Arquivos**: Lista todos os arquivos na pasta Downloads
3. **Filtragem**: Ignora diretórios e arquivos temporários (como `.zapzap_temp`)
4. **Categorização**: Identifica a extensão de cada arquivo
5. **Organização**: Move os arquivos para as pastas apropriadas
6. **Criação Automática**: Cria as pastas de destino se não existirem

## 🛡️ Segurança

- ✅ Não sobrescreve arquivos existentes
- ✅ Ignora pastas já criadas pelo programa
- ✅ Trata erros de forma adequada
- ✅ Funciona em múltiplas plataformas

## 📝 Código Exemplo

```go
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
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para:

1. Fazer um fork do projeto
2. Criar uma branch para sua feature (`git checkout -b feature/NovaFuncionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/NovaFuncionalidade`)
5. Abrir um Pull Request

## 📌 Melhorias Futuras

- [ ] Adicionar mais tipos de arquivo
- [ ] Permitir configuração customizada de categorias
- [ ] Adicionar interface gráfica
- [ ] Suporte para múltiplas pastas
- [ ] Sistema de logs detalhado
- [ ] Desfazer última organização
- [ ] Agendamento automático (cron job)

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👤 Autor

**MarkHiarley**

- GitHub: [@MarkHiarley](https://github.com/MarkHiarley)

## 🐛 Reportar Bugs

Encontrou um bug? Abra uma [issue](https://github.com/MarkHiarley/organizador-de-pasta/issues) descrevendo o problema.

---

⭐ Se este projeto foi útil, considere dar uma estrela!
