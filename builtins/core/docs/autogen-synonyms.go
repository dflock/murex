package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!and`:            `and`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`if`:              `if`,
	`post`:            `post`,
	`brace-quote`:     `brace-quote`,
	`pt`:              `pt`,
	`out`:             `out`,
	`tout`:            `tout`,
	`read`:            `read`,
	`alter`:           `alter`,
	`murex-docs`:      `murex-docs`,
	`err`:             `err`,
	`getfile`:         `getfile`,
	`try`:             `try`,
	`export`:          `export`,
	`f`:               `f`,
	`catch`:           `catch`,
	`or`:              `or`,
	`set`:             `set`,
	`prepend`:         `prepend`,
	`get`:             `get`,
	`g`:               `g`,
	`swivel-table`:    `swivel-table`,
	`>>`:              `>>`,
	`event`:           `event`,
	`rx`:              `rx`,
	`tread`:           `tread`,
	`and`:             `and`,
	`global`:          `global`,
	`append`:          `append`,
	`>`:               `>`,
	`ttyfd`:           `ttyfd`,
	`swivel-datatype`: `swivel-datatype`,
	`trypipe`:         `trypipe`,
}
