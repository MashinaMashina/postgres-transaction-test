## Тестирование работы Postgres при разных уровнях изоляции транзакций

- тестирование работы последовательностей при не закоммиченных добавлениях записей
- тестирование грязного чтения, неповторяемого чтения, фантомного чтения и аноманий сериализации.

Описание работы различных режимов: https://postgrespro.ru/docs/postgrespro/9.5/transaction-iso