package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!global`:         `global`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!or`:             `or`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!and`:            `and`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`tout`:            `tout`,
	`or`:              `or`,
	`alter`:           `alter`,
	`getfile`:         `getfile`,
	`pt`:              `pt`,
	`read`:            `read`,
	`try`:             `try`,
	`out`:             `out`,
	`get`:             `get`,
	`trypipe`:         `trypipe`,
	`export`:          `export`,
	`prepend`:         `prepend`,
	`brace-quote`:     `brace-quote`,
	`>>`:              `>>`,
	`event`:           `event`,
	`swivel-table`:    `swivel-table`,
	`post`:            `post`,
	`err`:             `err`,
	`rx`:              `rx`,
	`catch`:           `catch`,
	`swivel-datatype`: `swivel-datatype`,
	`>`:               `>`,
	`ttyfd`:           `ttyfd`,
	`f`:               `f`,
	`g`:               `g`,
	`and`:             `and`,
	`set`:             `set`,
	`append`:          `append`,
	`murex-docs`:      `murex-docs`,
	`tread`:           `tread`,
	`if`:              `if`,
	`global`:          `global`,
}
