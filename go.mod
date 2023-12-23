module example.com/gomod2nix-template

go 1.21.5

require example.com/greet v0.0.0-00010101000000-000000000000

replace example.com/greet => ./greet
