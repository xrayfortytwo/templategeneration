grammar Calc ;

@header {
    // calc is for calculating
}

calc : expre* EOF ;
expre : num? '*' num '=' num ;
num : NUMBER ;
NUMBER : [0-9]+ ;

