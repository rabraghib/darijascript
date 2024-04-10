import { styleTags, tags as t } from '@lezer/highlight';
import { LRParser } from '@lezer/lr';
import {
  foldNodeProp,
  foldInside,
  indentNodeProp,
  delimitedIndent,
} from '@codemirror/language';

const parser = getParser();

export const parserWithMetadata = parser.configure({
  props: [
    styleTags({
      '9ayed': t.definitionKeyword,
      'ilakan sinn ilamakanch ma7dBa9i rjje3 9ayed WAA fonksyon':
        t.controlKeyword,
      'True False s7i7 S7I7 ghalt GHALT': t.bool,
      'CallExpression/FieldExpression/FieldIdentifier': t.function(
        t.propertyName
      ),
      StatementIdentifier: t.labelName,
      Identifier: t.variableName,
      'CallExpression/Identifier': t.function(t.variableName),
      'FunctionDeclarator/Identifier FunctionDeclarator/DestructorName':
        t.function(t.definition(t.variableName)),
      NamespaceIdentifier: t.namespace,
      OperatorName: t.operator,
      ArithOp: t.arithmeticOperator,
      LogicOp: t.logicOperator,
      BitOp: t.bitwiseOperator,
      CompareOp: t.compareOperator,
      AssignOp: t.definitionOperator,
      UpdateOp: t.updateOperator,
      LineComment: t.lineComment,
      BlockComment: t.blockComment,
      Number: t.number,
      String: t.string,
      '( )': t.paren,
      '[ ]': t.squareBracket,
      '{ }': t.brace,
      '.': t.derefOperator,
      ', ;': t.separator,
    }),
    indentNodeProp.add({
      Application: delimitedIndent({ closing: ')', align: false }),
    }),
    foldNodeProp.add({
      Application: foldInside,
    }),
  ],
});

function getParser() {
  const spec_identifier: any = {
    __proto__: null,
    true: 24,
    false: 24,
    s7i7: 24,
    S7I7: 24,
    ghalt: 24,
    GHALT: 24,
    ilakan: 72,
    sinn: 74,
    ma7dBa9i: 78,
    rjje3: 88,
    '9ayed': 92,
    WAA: 96,
    fonksyon: 104,
  };
  return LRParser.deserialize({
    version: 14,
    states:
      "+|Q]QPOOOOQO'#Dm'#DmOOQO'#Cc'#CcOOQO'#Dp'#DpO!gQPO'#ClO$YQQO'#CnO$aQPO'#CmOOQO'#Do'#DoO#XQQO'#DoO$fQQO'#DkO&RQQO'#CaO!gQPO'#C|OOQO'#Dk'#DkO&YQPO'#DWOOQO'#Dj'#DjOOQO'#Dd'#DdQ]QPOOO&aQPO'#DPO&aQPO'#DSO&fQPO'#DXO!_QPO'#DZO&mQPO'#D]O!_QPO'#DaO&xQQO,59WO!gQPO,58|O'PQPO'#CoOOQO,59X,59XO!gQPO,59POOQO'#Cu'#CuOOQO'#Cw'#CwO!gQPO,59_O!gQPO,59_O!gQPO,59_O!gQPO,59_O!gQPO,59_OOQO,58{,58{OOQO,59h,59hOOQO,59r,59rO'WQPO,59rOOQO-E7b-E7bO]QPO,59kO]QPO,59nO'_QQO,59sOOQO,59s,59sO'fQPO,59uOOQO'#D`'#D`O'kQPO,59wO'pQPO,59yO'uQPO,59{OOQO1G.r1G.rO'zQQO1G.hO([QQO'#DqO(fQPO,59ZO(kQQO1G.kOOQO1G.y1G.yO)jQQO1G.yO*cQQO1G.yO*pQQO1G.yO*wQQO1G.yOOQO1G/^1G/^O+rQPO1G/VOOQO1G/Y1G/YOOQO1G/_1G/_OOQO1G/a1G/aOOQO1G/c1G/cOOQO1G/e1G/eO&pQPO'#DcO]QPO1G/gO!gQPO'#DeO-VQPO,5:]OOQO1G.u1G.uOOQO7+$V7+$VO]QPO7+$qO-_QPO'#DuO-gQPO,59}OOQO7+%R7+%RO-lQQO,5:POOQO-E7c-E7cOOQO<<H]<<H]O&pQPO'#DfO-vQPO,5:aOOQO1G/i1G/iOOQO,5:Q,5:QOOQO-E7d-E7d",
    stateData:
      '.V~OPOSQOS!]OS~OS^OYROZRO[RO]RO_SOnZOqZOrZOtaOwbOy]O|cO!OdO!QeO!UfO!`QO!bPO~OYROZRO[RO]RO_SOnZOqZOrZO!`QO!bPO~OfhOS!cXe!cXh!cXj!cXl!cXm!cXn!cXo!cX!f!cX!g!cX!h!cX^!cXd!cXW!cX~O_bX~P#XO_iO~OekOS!_Xh!_Xj!_Xl!_Xm!_Xn!_Xo!_X!f!_X!g!_X!h!_X^!_Xd!_XW!_X~OhnOjoOlpOmpOnqOorO!flO!glO!hmO~OSsO~P%dOxuO~P]O_SO~OS{O~P!gO]!OO!`}O!b}O~O^!RO~P%dO^!eP~P!gOx!]O~P]OS!`O~P%dOS!aO~OS!bO~OS!cO~O_!dO~OSUi^UidUiWUi~P%dOd!fO^!eX~P%dO^!hO~OW!iO~P%dOhnOorO!flO!glOSgijgilgimgi!hgi^gidgiWgi~OnqO~P(rOhnOorO!flO!glOSgilgimgi^gidgiWgi~OjoOnqO!hmO~P)qOngi~P(rOhnO!flO!glOSgijgilgimgingiogi!hgi^gidgiWgi~Ou!jOSsiYsiZsi[si]si_sinsiqsirsitsiwsiysi|si!Osi!Qsi!Usi!Zsi!`si!bsixsi~Od!fO^!ea~Od!qO^!iX~O^!sO~O^!Xad!Xa~P%dOd!qO^!ia~OZQYPoP~',
    goto: ")s!jPPPPP!k!u#]P#yPPPPPP$a$}%e%{PPP!uP&OP&]PPPP!uPP!kPP!kPPP!k!kP!kP!kP!k&g!kP&p&s&}'TPPP'Z'mP(mP)V$})mPPP)pa^O]`vxy!e!jy[OSZ]`chiknopqrvxy!e!f!jxTOSZ]`chiknopqrvxy!e!f!jQ|dR!QfyWOSZ]`chiknopqrvxy!e!f!jxVOSZ]`chiknopqrvxy!e!f!jQxaRybyVOSZ]`chiknopqrvxy!e!f!jyUOSZ]`chiknopqrvxy!e!f!jRjUgnYgz!S!T!V!X!Y!Z![!naoYgz!S!T!V!Y!nQ!PeQ!k!dR!t!qR!e!QQ`OQv]Tw`vQ!g!TR!o!gQ!r!kR!u!rW_O]`vQ!^xQ!_yQ!m!eR!p!j`YO]`vxy!e!jQgSQtZQzcQ!ShQ!TiQ!VkQ!WnQ!XoQ!YpQ!ZqQ![rR!n!f}QOSZ]`cdfhiknopqrvxy!e!f!jyXOSZ]`chiknopqrvxy!e!f!jR!UiR!l!d",
    nodeNames:
      '⚠ LineComment BlockComment Program ; ExpressionStatement AssignmentExpression Identifier ] ArrayAccess IntegerLiteral FloatingPointLiteral BooleanLiteral StringLiteral ) ( ParenthesizedExpression MethodInvocation MethodName ArgumentList , [ AssignOp BinaryExpression CompareOp CompareOp LogicOp BitOp BitOp LogicOp ArithOp ArithOp UnaryExpression LogicOp BitOp IfStatement ilakan sinn ForStatement ma7dBa9i } { Block ReturnStatement rjje3 VariableDeclaration 9ayed ThrowStatement WAA Throws Definition MethodDeclaration fonksyon InferredParameters',
    maxTerm: 71,
    nodeProps: [
      ['isolate', -3, 1, 2, 13, ''],
      [
        'group',
        -10,
        4,
        5,
        35,
        38,
        42,
        43,
        45,
        47,
        49,
        51,
        'Statement',
        -11,
        6,
        7,
        9,
        10,
        11,
        12,
        13,
        16,
        17,
        23,
        32,
        'Expression',
      ],
      ['openedBy', 14, '(', 40, '{'],
      ['closedBy', 15, ')', 41, '}'],
    ],
    skippedNodes: [0, 1, 2],
    repeatNodeCount: 3,
    tokenData:
      ",X~RvX^#ipq#iqr$^rs$kst%utu&^vw&rxy'Pyz'Uz{'Z{|'`|}'e}!O'`!O!P'j!P!Q(U!Q![)h!]!^*X!^!_*^!_!`*f!`!a*n!c!}*v!}#O+[#P#Q+a#Q#R+f#R#S&^#T#o&^#o#p+k#p#q+p#q#r+}#r#s,S#y#z#i$f$g#i#BY#BZ#i$IS$I_#i$I|$JO#i$JT$JU#i$KV$KW#i&FU&FV#i~#nY!]~X^#ipq#i#y#z#i$f$g#i#BY#BZ#i$IS$I_#i$I|$JO#i$JT$JU#i$KV$KW#i&FU&FV#iR$cPqP!_!`$fQ$kOhQ~$nWOY$kZr$krs%Ws#O$k#O#P%]#P;'S$k;'S;=`%o<%lO$k~%]O]~~%`TOY$kYZ$kZ;'S$k;'S;=`%o<%lO$k~%rP;=`<%l$k~%zSP~OY%uZ;'S%u;'S;=`&W<%lO%u~&ZP;=`<%l%u~&cT!`~tu&^!Q![&^!c!}&^#R#S&^#T#o&^~&wP!h~vw&z~'POj~~'UO_~~'ZO^~~'`Oo~~'eOn~~'jOd~~'mP!Q!['p~'uQZ~!Q!['p#R#S'{~(OQ!Q!['p#R#S'{~(ZQo~z{(a!P!Q%u~(dTOz(az{(s{;'S(a;'S;=`)b<%lO(a~(vVOz(az{(s{!P(a!P!Q)]!Q;'S(a;'S;=`)b<%lO(a~)bOQ~~)eP;=`<%l(a~)mRY~!O!P)v!Q![)h#R#S*O~){PZ~!Q!['p~*RQ!Q![)h#R#S*O~*^OS~~*cP!f~!_!`$f~*kPf~!_!`$f~*sP!g~!_!`$f~*{T!b~tu*v!Q![*v!c!}*v#R#S*v#T#o*v~+aOe~~+fOW~~+kOl~~+pOy~~+uPl~#p#q+x~+}Om~~,SOx~~,XOr~",
    tokenizers: [0, 1],
    topRules: { Program: [0, 3] },
    dynamicPrecedences: { '63': -1 },
    specialized: [
      { term: 62, get: (value: any) => spec_identifier[value] || -1 },
    ],
    tokenPrec: 598,
  });
}
