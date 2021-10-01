# userCrud
Небольшой gRPC CRUD сервис для операций с пользователями

При написании использовались принципы чистой архитектуры. Слой бизнес сущностей и бизнес логики не зависит ни от одного из друх, 
что позволяет просто заменять компоненты и добавлять новые реализации. Слой repository содержит реализацию работы с объектом на уровне базы данных.
Слой интерфейсов содержит в себе обработчики запросов.

В качестве базы данных реализована MongoDB

