{ pkgs, config, lib, ... }:

{
  dotenv.enable = true;

  languages.go = {
    enable = true;
    package = pkgs.go_1_25;
  };

  packages = with pkgs; [
    gnumake
    gcc
    google-wire
    golangci-lint
    postgresql_16
    redis
  ];

  services.postgres = {
    enable = true;
    package = pkgs.postgresql_16;

    # Matches your .env DB_NAME
    initialDatabases = [{ name = "goth"; }];

    # Matches your .env DB_USER and DB_PASSWORD
    initialScript = ''
      CREATE USER postgres WITH SUPERUSER PASSWORD 'postgres';
      ALTER DATABASE goth OWNER TO postgres;
    '';

    listen_addresses = "127.0.0.1";
    port = 5432;
  };

  services.redis = {
    enable = true;
    port = 6379;
  };

  scripts.gen-wire.exec = "wire ./cmd/app";

  processes.goat-api.exec = "make run";

  # 8. Shell Hook
  enterShell = ''
    echo "GOAT API Development Environment Loaded!"
    echo "-----------------------------------------"
    echo "DATABASE: postgres://postgres:postgres@localhost:5432/goth"
    echo "REDIS:    localhost:6379"
    echo "GO:       $(go version)"
    echo "-----------------------------------------"
  '';
}
