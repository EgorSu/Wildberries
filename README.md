# Wildberries

## L0:
Сервис находится в папке service(main.go). 
Publisher читает файлы из рublisher/data и записывает из в канал nats-streaming. Параметры бд и nats-streaming указаны в service/main.go.
## L1:
В некоторых заданиях используются флаги для передачи параметров:  

task4:  
        -num, кол-во воркеров, defaultValue = 1  
task8: 	  
        -val, число для изменения, defaultValue = 1  
	-bitNum, номер бита, который будем менять, defaultValue = 0   
	-bitVal, новое значение бита, defaultValue = 0   
task19:   
        -word, слово для переворачивания, defaultValue = "главрыба"  
task26:   
        -word, слово для проверки, defaultValue = "abcd"   
task20:  
	Строку для изменения порядка слов можно подать на вход как аргументы командной строки.  
   
