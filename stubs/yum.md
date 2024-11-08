## <NAME>

## Установка
```bash
# Установка yum-config-manager если отсутствует
sudo yum install yum-utils

# Добавление зеркала в список репозиториев
sudo yum-config-manager --add-repo <REPO>

# Установка TailScale
sudo yum install tailscale

# Включение и запуск tailscaled
sudo systemctl enable --now tailscaled

# Запуск TailScale
sudo tailscale up
```