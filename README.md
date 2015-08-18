# matic
Need a hand with dat quick arithmetic? Matic is here to save your ass. Whenever you open up a new terminal you will be prompted to answer a randomly generated sum that you *should* be able to solve with simple mental arithmetic.

## Installation
To install matic you can either build from source (like a normal go project) through using `go get` etc, or download the latest stable version in the [releases page](https://github.com/paked/matic/releases).

Once you have gotten hold of a build, you should append this line to either your `.bashrc` or `.zshrc`: `./path/to/matic`.

## Options
By default, matic will test you on addition and subtraction sums where the numbers will be `>= 0` and `<= 12`. You can change this by specifiying a flag or environment variable as shown below

Flag Name | Environment Variable Name | Result                   | Default
----------|---------------------------|--------------------------|
--min | MIN | Change the minimum value | 0
--max | MAX | Change the maximum value | 12
--addition | ADDITION | Whether to enable additon | true 
--subtraction | SUBTRACTION | Whether to enable subtraction | true
--multiplication | MULTIPLICATION | Whether to enable multiplication | false
--division | DIVISION | Whether to enable division | false
