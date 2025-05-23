## FEDORA

## Установка
```bash
# Добавление зеркала в список репозиториев
sudo dnf config-manager --add-repo https://raw.githubusercontent.com/hopleus/tailscale-mirror/main/data/stable/fedora/tailscale.repo

# Установка TailScale
sudo dnf install tailscale

# Включение и запуск tailscaled
sudo systemctl enable --now tailscaled

# Запуск TailScale
sudo tailscale up
```