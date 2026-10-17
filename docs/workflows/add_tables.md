DriveType
-> id, Code
FuelType
-> id, Code
TransmissionType
-> id, Code
+PowerHp
+-> field Car
+BodyType
+-> id, Code
Engine - пока не делаем!

В модели нужно добавить Тип привода в Name

на потом

Relation("BodyType").
Relation("Brand").
Relation("Color").
Relation("DriveType").
Relation("FuelType").
Relation("Model").
Relation("Model.Brand").
Relation("TransmissionType").

```go
type CreateCarRequest struct {
	BodyType         *CodeReference `json:"body_type,omitempty"`
	Brand            *NameReference `json:"brand,omitempty"`
	Color            *CarColor      `json:"color,omitempty"`
	Description      *string        `json:"description,omitempty"`
	DriveType        *CodeReference `json:"drive_type,omitempty"`
	FuelType         *CodeReference `json:"fuel_type,omitempty"`
	Model            *CarModel      `json:"model,omitempty"`
	Name             string         `json:"name"`
	PowerHp          *string        `json:"power_hp,omitempty"`
	TransmissionType *CodeReference `json:"transmission_type,omitempty"`
}

```

```bash
curl -X POST http://localhost:8080/cars \
-H "Content-Type: application/json" \
-d '{
  "name":"Aurus Senat Limousine",
  "description":"Full relations t_e_s_t 2",
  "power_hp": "500",

  "color":{
    "code":"black"
  },

  "brand":{
    "name":"Aurus"
  },

  "model":{
    "name":"Senat"
  },

  "body_type":{
    "code":"limousine"
  },

  "fuel_type":{
    "code":"hybrid"
  },

  "transmission_type":{
    "code":"automatic"
  },

  "drive_type":{
    "code":"awd"
  }
}'
```

```bash
curl -X PATCH http://localhost:8080/cars/0281aeff-1586-43b4-89e6-0c06b10d5030 \
-H "Content-Type: application/json" \
-d '{
  "name":"Aurus Senat Patriot+",
  "description":"PATCH relations te__st",
  "power_hp": "503",
    "brand":{
    "name":"Aurus+"
  },
  "model":{
    "name":"Senat+"
  },
  "color":{
    "code":"white+"
  },
  "body_type":{
    "code":"sedan+"
  },
  "fuel_type":{
    "code":"plug-in-hybrid+"
  },
  "transmission_type":{
    "code":"automatic+"
  },
  "drive_type":{
    "code":"awd+"
  }
}'
```

```text
{
    "body_type":{"code":"sedan","id":"7efb078b-b6bb-4b13-9d57-efe7ab63a1c7"},
    "color":{"code":"white","id":"1b81750c-b0a7-4720-9d53-1b5f91345b08"},
    "created_at":"2026-10-17T20:09:50.828523Z",
    "description":"PATCH relations te__st",
    "drive_type":{"code":"awd","id":"2219d18d-e35c-486e-8a0f-a387c16acec7"},
    "id":"b88402b6-64ad-420d-a5e9-154d99c0583f",
    "name":"Aurus Senat Patriot",
    "transmission_type":{"code":"automatic","id":"f5fd5f78-a112-43aa-92bc-ccb808d23c8f"},
    "updated_at":"2026-10-17T20:16:47.336896Z"
    }
```

```text
{
    "body_type":{"code":"sedan","id":"7efb078b-b6bb-4b13-9d57-efe7ab63a1c7"},
    "color":{"code":"white","id":"1b81750c-b0a7-4720-9d53-1b5f91345b08"},
    "drive_type":{"code":"awd","id":"2219d18d-e35c-486e-8a0f-a387c16acec7"},
    "transmission_type":{"code":"automatic","id":"f5fd5f78-a112-43aa-92bc-ccb808d23c8f"},
    }
```
