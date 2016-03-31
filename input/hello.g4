grammar Hello ;

subscript : descript* EOF ;
descript : 'surname' '=' surname ',' (id '.' company '@' email)? ;
company : NAME ;
email : NAME
name : NAME ;
surname : NAME ;
NAME : [a-z]+ ;
