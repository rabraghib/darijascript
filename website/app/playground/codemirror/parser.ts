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
    ilamakanch: 76,
    ma7dBa9i: 80,
    rjje3: 90,
    '9ayed': 94,
    WAA: 98,
    fonksyon: 106,
  };
  return LRParser.deserialize({
    version: 14,
    states:
      "+|Q]QPOOOOQO'#Dn'#DnOOQO'#Cc'#CcOOQO'#Dq'#DqO!jQPO'#ClO$]QQO'#CnO$dQPO'#CmOOQO'#Dp'#DpO#[QQO'#DpO$iQQO'#DlO&UQQO'#CaO!jQPO'#C|OOQO'#Dl'#DlO&]QPO'#DXOOQO'#Dk'#DkOOQO'#De'#DeQ]QPOOO&dQPO'#DPO&dQPO'#DTO&iQPO'#DYO!bQPO'#D[O&pQPO'#D^O!bQPO'#DbO&{QQO,59WO!jQPO,58|O'SQPO'#CoOOQO,59X,59XO!jQPO,59POOQO'#Cu'#CuOOQO'#Cw'#CwO!jQPO,59_O!jQPO,59_O!jQPO,59_O!jQPO,59_O!jQPO,59_OOQO,58{,58{OOQO,59h,59hOOQO,59s,59sO'ZQPO,59sOOQO-E7c-E7cO]QPO,59kO]QPO,59oO'bQQO,59tOOQO,59t,59tO'iQPO,59vOOQO'#Da'#DaO'nQPO,59xO'sQPO,59zO'xQPO,59|OOQO1G.r1G.rO'}QQO1G.hO(_QQO'#DrO(iQPO,59ZO(nQQO1G.kOOQO1G.y1G.yO)mQQO1G.yO*fQQO1G.yO*sQQO1G.yO*zQQO1G.yOOQO1G/_1G/_O+uQPO1G/VOOQO1G/Z1G/ZOOQO1G/`1G/`OOQO1G/b1G/bOOQO1G/d1G/dOOQO1G/f1G/fO&sQPO'#DdO]QPO1G/hO!jQPO'#DfO-]QPO,5:^OOQO1G.u1G.uOOQO7+$V7+$VO]QPO7+$qO-eQPO'#DvO-mQPO,5:OOOQO7+%S7+%SO-rQQO,5:QOOQO-E7d-E7dOOQO<<H]<<H]O&sQPO'#DgO-|QPO,5:bOOQO1G/j1G/jOOQO,5:R,5:ROOQO-E7e-E7e",
    stateData:
      '.]~OPOSQOS!^OS~OS^OYROZRO[RO]RO_SOnZOqZOrZOtaOvaOxbOz]O}cO!PdO!ReO!VfO!aQO!cPO~OYROZRO[RO]RO_SOnZOqZOrZO!aQO!cPO~OfhOS!dXe!dXh!dXj!dXl!dXm!dXn!dXo!dX!g!dX!h!dX!i!dX^!dXd!dXW!dX~O_bX~P#[O_iO~OekOS!`Xh!`Xj!`Xl!`Xm!`Xn!`Xo!`X!g!`X!h!`X!i!`X^!`Xd!`XW!`X~OhnOjoOlpOmpOnqOorO!glO!hlO!imO~OSsO~P%gOyuO~P]O_SO~OS{O~P!jO]!OO!a}O!c}O~O^!RO~P%gO^!fP~P!jOy!]O~P]OS!`O~P%gOS!aO~OS!bO~OS!cO~O_!dO~OSUi^UidUiWUi~P%gOd!fO^!fX~P%gO^!hO~OW!iO~P%gOhnOorO!glO!hlOSgijgilgimgi!igi^gidgiWgi~OnqO~P(uOhnOorO!glO!hlOSgilgimgi^gidgiWgi~OjoOnqO!imO~P)tOngi~P(uOhnO!glO!hlOSgijgilgimgingiogi!igi^gidgiWgi~Ou!jOSsiYsiZsi[si]si_sinsiqsirsitsivsixsizsi}si!Psi!Rsi!Vsi![si!asi!csiysi~Od!fO^!fa~Od!qO^!jX~O^!sO~O^!Yad!Ya~P%gOd!qO^!ja~OZQYPoP~',
    goto: ")t!kPPPPP!l!v#^P#zPPPPPP$b%O%f%|PPP!vP&PP&^PPPP!vPP!lPPP!lPPP!l!lP!lP!lP!l&h!lP&q&t'O'UPPP'['nP(nP)W%O)nPPP)qa^O]`vxy!e!jy[OSZ]`chiknopqrvxy!e!f!jxTOSZ]`chiknopqrvxy!e!f!jQ|dR!QfyWOSZ]`chiknopqrvxy!e!f!jxVOSZ]`chiknopqrvxy!e!f!jQxaRybyVOSZ]`chiknopqrvxy!e!f!jyUOSZ]`chiknopqrvxy!e!f!jRjUgnYgz!S!T!V!X!Y!Z![!naoYgz!S!T!V!Y!nQ!PeQ!k!dR!t!qR!e!QQ`OQv]Tw`vQ!g!TR!o!gQ!r!kR!u!rW_O]`vQ!^xQ!_yQ!m!eR!p!j`YO]`vxy!e!jQgSQtZQzcQ!ShQ!TiQ!VkQ!WnQ!XoQ!YpQ!ZqQ![rR!n!f}QOSZ]`cdfhiknopqrvxy!e!f!jyXOSZ]`chiknopqrvxy!e!f!jR!UiR!l!d",
    nodeNames:
      '⚠ LineComment BlockComment Program ; ExpressionStatement AssignmentExpression Identifier ] ArrayAccess IntegerLiteral FloatingPointLiteral BooleanLiteral StringLiteral ) ( ParenthesizedExpression MethodInvocation MethodName ArgumentList , [ AssignOp BinaryExpression CompareOp CompareOp LogicOp BitOp BitOp LogicOp ArithOp ArithOp UnaryExpression LogicOp BitOp IfStatement ilakan sinn ilamakanch ForStatement ma7dBa9i } { Block ReturnStatement rjje3 VariableDeclaration 9ayed ThrowStatement WAA Throws Definition MethodDeclaration fonksyon InferredParameters',
    maxTerm: 72,
    nodeProps: [
      ['isolate', -3, 1, 2, 13, ''],
      [
        'group',
        -10,
        4,
        5,
        35,
        39,
        43,
        44,
        46,
        48,
        50,
        52,
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
      ['openedBy', 14, '(', 41, '{'],
      ['closedBy', 15, ')', 42, '}'],
    ],
    skippedNodes: [0, 1, 2],
    repeatNodeCount: 3,
    tokenData:
      ",X~RvX^#ipq#iqr$^rs$kst%utu&^vw&rxy'Pyz'Uz{'Z{|'`|}'e}!O'`!O!P'j!P!Q(U!Q![)h!]!^*X!^!_*^!_!`*f!`!a*n!c!}*v!}#O+[#P#Q+a#Q#R+f#R#S&^#T#o&^#o#p+k#p#q+p#q#r+}#r#s,S#y#z#i$f$g#i#BY#BZ#i$IS$I_#i$I|$JO#i$JT$JU#i$KV$KW#i&FU&FV#i~#nY!^~X^#ipq#i#y#z#i$f$g#i#BY#BZ#i$IS$I_#i$I|$JO#i$JT$JU#i$KV$KW#i&FU&FV#iR$cPqP!_!`$fQ$kOhQ~$nWOY$kZr$krs%Ws#O$k#O#P%]#P;'S$k;'S;=`%o<%lO$k~%]O]~~%`TOY$kYZ$kZ;'S$k;'S;=`%o<%lO$k~%rP;=`<%l$k~%zSP~OY%uZ;'S%u;'S;=`&W<%lO%u~&ZP;=`<%l%u~&cT!a~tu&^!Q![&^!c!}&^#R#S&^#T#o&^~&wP!i~vw&z~'POj~~'UO_~~'ZO^~~'`Oo~~'eOn~~'jOd~~'mP!Q!['p~'uQZ~!Q!['p#R#S'{~(OQ!Q!['p#R#S'{~(ZQo~z{(a!P!Q%u~(dTOz(az{(s{;'S(a;'S;=`)b<%lO(a~(vVOz(az{(s{!P(a!P!Q)]!Q;'S(a;'S;=`)b<%lO(a~)bOQ~~)eP;=`<%l(a~)mRY~!O!P)v!Q![)h#R#S*O~){PZ~!Q!['p~*RQ!Q![)h#R#S*O~*^OS~~*cP!g~!_!`$f~*kPf~!_!`$f~*sP!h~!_!`$f~*{T!c~tu*v!Q![*v!c!}*v#R#S*v#T#o*v~+aOe~~+fOW~~+kOl~~+pOz~~+uPl~#p#q+x~+}Om~~,SOy~~,XOr~",
    tokenizers: [0, 1],
    topRules: { Program: [0, 3] },
    dynamicPrecedences: { '64': -1 },
    specialized: [
      { term: 63, get: (value: any) => spec_identifier[value] || -1 },
    ],
    tokenPrec: 604,
  });
}
