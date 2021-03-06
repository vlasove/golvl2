### 3. Посетитель
**Посетитель** - это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

**Посетитель** - это способ добавления нового функционала существующим классам (объектам) с минимумом изменений исходных классов. Добавление нового функционала сводится к реализации класса **посетителя** и, возможно, небольшого изменения исходных классов, направленных на взаимодействие с посетитлем.

**Аналогия из жизни** - страховой агент, который `посещает` разные организации. Для каждой организации `посетитель` приготовил специальное предложение. (у каждого типа посетитель выдает новый функционал, необходимый для этого типа).

#### 3.1 Применимость
* Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
* Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции, но вы не хотите «засорять» классы такими операциями.
* Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

#### 3.2 Преимущества
* Упрощает добавление операций, работающих со сложными структурами объектов.
* Объединяет родственные операции в одном классе.
* Посетитель может накапливать состояние при обходе структуры элементов.
#### 3.3 Недостатки
* Паттерн не оправдан, если иерархия элементов часто меняется.
* Может привести к нарушению инкапсуляции элементов (если посетителям нужен доступ к приватным полям).