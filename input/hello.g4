grammar Hello ;

subscript : descript* EOF ;
descript : 'surname' '=' surname ',' (surname '@' email)? ;
company : NAME ;
email : NAME
name : NAME ;
surname : NAME ;
