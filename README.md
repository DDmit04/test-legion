### Тестовое задание - [постановка](./Тестовое%20задание%20для%20Golang-разработчика_Легион.docx)

## Использование пакета

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DDmit04/test-legion/src"
	"github.com/DDmit04/test-legion/src/models/input"
)

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//создание и валидация данных портов и хостов
	hosts, err := input.CreateHostsListInput([]string{"yandex.ru", "google.com"})
	if err != nil {
		// ошибки валидации данных
		panic(err)
	}

	ports, err := input.CreatePortsListInput([]int{80, 21})
	if err != nil {
		// ошибки валидации данных
		panic(err)
	}

	// количество одновременных запросов
	connections := 10
	// время, по истичению которого порт будет отмечен как неответевший
	timeout := 1 * time.Second
	scn := src.NewScanner(connections, timeout)
	// функция сканирования - возвращает буферизированный канал, в который пишутся результаты
	// как только все порты будут просканированы или сканирование будет прервано контекстом, канал будет закрыт
	channel, err := scn.Scan(ctx, hosts, ports)
	if err != nil {
		// ошибки чтения данных
		panic(err)
	}

	for val := range channel {
		fmt.Println(val)
	}
}

```

### Данные, доступные для ввода
- `input.CreatePortsListInput` - принимает список портов в виде `[]int` 
- `input.CreateHostsListInput` - принимает список хостов в виде `[]string`
- `input.CreatePortsRangeInput` - принимает диопазон портов в виде `string`

### Возможные ошибки
- `reader not found` - не удалось прочитать данные портов или хостов
- `invalid input` - общая ошибка для неверного ввода
- `invalid host` - введённый хост не является валидным
- `invalid ip` - введённый IP не является валидным
- `host ips read error` - не удалось получить список IP, которые принадлежат хосту
- `unexpected ports range format` - невалидный формат диопазона портов
- `invalid port value` - невалидный порт
- `zero ports not accepted` - нельзя указывать порт 0
- `negative ports not accepted` - нельзя указывать отрицательное число в качестве порта
- `ports more than 65535 not accepted` - нельзя указывать порт больше 65535


### Формат данных в результате
- `Host` (string) — доменное имя или хост, по которому выполнялось сканирование
- `IP` (net.IP) — разрешённый IP‑адрес цели (IPv4/IPv6)
- `Port` (int) — номер проверяемого порта (например, 80, 443, 8080)
- `State` (string) — статус порта после проверки. Возможные значения: open, closed, timeout, unreachable, canceled, error
- `Duration` (time.Duration) — время, затраченное на проверку порта (от начала попытки соединения до получения результата)
- `Err` (error) — исходная ошибка, если проверка завершилась нештатно
