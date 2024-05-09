## DEBIAN (TRIXIE)

## Установка
```bash
# Добавление TailScale GPG ключа
sudo mkdir -p --mode=0755 /usr/share/keyrings
curl -fsSL https://raw.githubusercontent.com/hopleus/TailScale-APT-Mirror/main/data/stable/debian/trixie.noarmor.gpg | sudo tee /usr/share/keyrings/tailscale-mirror-archive-keyring.gpg > /dev/null

# Добавление зеркала в источники APT
curl -fsSL https://raw.githubusercontent.com/hopleus/TailScale-APT-Mirror/main/data/stable/debian/trixie.tailscale-keyring.list | sudo tee /etc/apt/sources.list.d/tailscale-mirror.list

# Установка TailScale
sudo apt-get update && sudo apt-get install tailscale

# Запуск TailScale
sudo tailscale up
```