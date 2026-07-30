package app

const defaultConfigTemplate = `server:
  address: ":8080" # port on which we are listening during gamesync serve

remote:
  address: "127.0.0.1:8080" # address of the server that the client accesses during synchronization

profiles:
  # example Factorio profile
  - name: factorio
    game: factorio
    root: C:\Path\To\Factorio

  # example Minecraft profile
  # - name: Minecraft
  #   game: minecraft
  #   root: C:\Users\USER\AppData\Roaming\.minecraft
`
