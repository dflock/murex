package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`(`:               `brace-quote`,
	`!and`:            `and`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`!set`:            `set`,
	`!event`:          `event`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`swivel-datatype`: `swivel-datatype`,
	`getfile`:         `getfile`,
	`tout`:            `tout`,
	`try`:             `try`,
	`prepend`:         `prepend`,
	`or`:              `or`,
	`alter`:           `alter`,
	`append`:          `append`,
	`swivel-table`:    `swivel-table`,
	`>`:               `>`,
	`and`:             `and`,
	`rx`:              `rx`,
	`tread`:           `tread`,
	`set`:             `set`,
	`event`:           `event`,
	`err`:             `err`,
	`out`:             `out`,
	`pt`:              `pt`,
	`if`:              `if`,
	`trypipe`:         `trypipe`,
	`murex-docs`:      `murex-docs`,
	`f`:               `f`,
	`catch`:           `catch`,
	`export`:          `export`,
	`global`:          `global`,
	`brace-quote`:     `brace-quote`,
	`>>`:              `>>`,
	`ttyfd`:           `ttyfd`,
	`g`:               `g`,
	`get`:             `get`,
	`post`:            `post`,
	`read`:            `read`,
}
