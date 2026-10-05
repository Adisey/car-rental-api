## команда создания пустой миграции:

./cmd/schema/atlas.sh migrate new truncate_cars \
--dir "file://db/migrations"

## Если осознанно менял миграции вручную, то нужно пересчитать контрольные суммы:

./cmd/schema/atlas.sh migrate hash --env local
