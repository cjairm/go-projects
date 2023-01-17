module create_module/hello

go 1.17

replace create_module/greetings => ../create_module

require create_module/greetings v0.0.0-00010101000000-000000000000
