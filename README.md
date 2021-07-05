# pharmacy_mdlp_2

pharmacy_mdlp_2 сервис для работы с сервисом МДЛП



## Реализованные методы

| Метод | Оперция                                        | Описание                                                     |
| ----- | ---------------------------------------------- | ------------------------------------------------------------ |
| POST  | /v2/documents/send                             | Метод отправки файла                                         |
| POST  | /v2/documents/send_large                       | Метод отправки файла большого размера                        |
| POST  | /v2/documents/cancel                           | Метод отмены отправки документа                              |
| POST  | /v2/documents/outcome                          | Метод получения списка исходящих документов                  |
| POST  | /v2/documents/income                           | Метод получения списка входящих документов                   |
| POST  | /v2/documents/income/mark_read                 | Метод передачи информации о прочтении документа              |
| GET   | /v2/documents/:docId                           | Метод получения метаданных документа                         |
| GET   | /v2/documents/request/:request_id              | Метод получения списка документов по ИД запроса              |
| GET   | /v2/documents/download/:docId                  | Метод получения документа по ИД                              |
| GET   | /v2/documents/:docId/ticket                    | Метод получения квитанции по ИД                              |
| POST  | /v2/reestr/sgtin/filter                        | Метод поиска по реестрам КИЗ                                 |
| POST  | /v2/reestr/sgtin/sgtins-by-list                | Метод поиска по реестрам КИЗ по списку значений              |
| POST  | /v2/sgtin/validate                             | Метод валидации кода маркировки                              |
| POST  | /v2/reestr/sgtin/public/sgtins-by-list         | Метод поиска по общедоступным реестрам КИЗ по списку значений |
| POST  | /v2/reestr/sgtin/public/archive/sgtins-by-list | Метод поиска по общедоступным реестрам КИЗ в архивном хранилище по списку значений |
| POST  | /v2/reestr/sscc/{sscc}/sgtins                  | Метод для получения информации о КИЗ в третичной упаковке    |
| POST  | /v2/reestr/trusted_partners/add                | Добавление доверенных контрагентов                           |
| POST  | /v2/reestr/trusted_partners/filter             | Метод фильтрации доверенных контрагентов                     |
| POST  | /v2/reestr/trusted_partners/delete             | Удаление доверенного контрагента                             |
| POST  | /v2/reestr_partners/filter                     | Метод фильтрации по субъектам обращения                      |
| GET   | /v2/reestr/sgtin/:sgtin                        | Метод для получения из реестров КИЗ детальной информации о КИЗ и связанным с ним ЛП |
| GET   | /v2/reestr/sgtin/archive/:sgtin                | Метод для получения детальной информации о КИЗ в архивном хранилище и связанным с ним ЛП |
| GET   | /v2/reestr/sscc/:sscc/hierarchy                | Метод для получения информации об иерархии вложенности третичной упаковки |
| GET   | /v2/reestr/sgtin/documents                     | Метод для получения перечня документов по идентификатору КИЗ |
| GET   | /v2/reestr/sscc/full-hierarchy                 | Метод получения полной иерархии третичной упаковки           |
| POST  | /v2/reestr/esklp/filter                        | Фильтрация по реестру ЕСКЛП                                  |


## Установка

- Развернуть Redis
- Настроить cryptoproCSP
- Настроить Stunnel

