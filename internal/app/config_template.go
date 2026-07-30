package app

const defaultConfigTemplate = `server:
  address: ":8080" # порт, на котором слушаем при gamesync serve

remote:
  address: "127.0.0.1:8080" # адрес сервера, к которому обращается клиент при синхронизации

profiles:
  # пример профиля Factorio
  - name: factorio
    game: factorio
    root: C:\Path\To\Factorio

  # пример профиля Minecraft
  # - name: Minecraft
  #   game: minecraft
  #   root: C:\Users\USER\AppData\Roaming\.minecraft
`
