### 2. Строитель
**Строитель** - это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. 

**Строитель** - это способ пошаговой сборки объекта известной сложности. Часто может существовать несколько строителей, которые выполняют свою работу (строят) по-разному. Для их управления может быть полезен **директор**. Строитель возвращает **продукт**.

**Продукт** - результат работы конкретного **строителя**.

**Директор** - объект, который заставляет работать строителей, но директор **НЕ ЗНАЕТ** про готовый продукт (не зависит от конкретной реализации строителя и продукта) 

**Строитель полезен** - когда итоговый объект требует достаточно обширной иницилазиации (замена телескопическому конструктору)

#### 2.1 Применимость
* Когда нужно избавиться от телескопического конструктора.
* Когда код должен создавать разные представления какого-то объекта. Например, внедорожник и спортивная машина.
* Когда нужно собирать сложные составные объекты (например рекурсивное создание деревьев)

#### 2.2 Преимущества
* Позволяет создавать продукты пошагово.
* Позволяет использовать один и тот же код для создания различных продуктов.
* Изолирует сложный код сборки продукта от его основной бизнес-логики.

#### 2.3 Недостатки
* Усложняет код программы из-за введения дополнительных классов.
* Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.