env "local" {
  src = "file://db/schema/schema.sql"

  dev = "postgres://app:app@192.168.100.100:5435/atlas_dev_db?sslmode=disable"

  url = "postgres://app:app@192.168.100.100:5434/car_rental?sslmode=disable"

  migration {
    dir = "file://db/migrations"
  }
}