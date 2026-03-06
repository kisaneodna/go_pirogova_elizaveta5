package actioninfo

import (
    "fmt"
)

type DataParser interface {
    Parse(datastring string) error
    ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
    for _, data := range dataset {
        if err := dp.Parse(data); err != nil {
            fmt.Printf("Ошибка парсинга %q: %v\n", data, err)
            continue
        }
        
        info, err := dp.ActionInfo()
        if err != nil {
            fmt.Printf("Ошибка получения информации для %q: %v\n", data, err)
            continue
        }
        
        fmt.Println(info)
    }
}
