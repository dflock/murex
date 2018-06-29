package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!or`:             `or`,
	`!if`:             `if`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`brace-quote`:     `brace-quote`,
	`tout`:            `tout`,
	`export`:          `export`,
	`append`:          `append`,
	`swivel-table`:    `swivel-table`,
	`swivel-datatype`: `swivel-datatype`,
	`tread`:           `tread`,
	`trypipe`:         `trypipe`,
	`event`:           `event`,
	`alter`:           `alter`,
	`prepend`:         `prepend`,
	`err`:             `err`,
	`set`:             `set`,
	`post`:            `post`,
	`pt`:              `pt`,
	`f`:               `f`,
	`murex-docs`:      `murex-docs`,
	`get`:             `get`,
	`global`:          `global`,
	`g`:               `g`,
	`read`:            `read`,
	`catch`:           `catch`,
	`getfile`:         `getfile`,
	`>`:               `>`,
	`try`:             `try`,
	`>>`:              `>>`,
	`and`:             `and`,
	`rx`:              `rx`,
	`or`:              `or`,
	`if`:              `if`,
	`out`:             `out`,
	`ttyfd`:           `ttyfd`,
}
