'use server';

export interface ICodingExample {
  name: string;
  code: string;
}

export async function getExamples(): Promise<ICodingExample[]> {
  return [
    {
      name: 'Hello, World!',
      code: '\ngolih("Salam 3likom!!");\n',
    },
    {
      name: 'FizzBuzz',
      code: '# FizzBuzz\n\nfizzBuzz(100);\n\nfonksyon fizzBuzz(n) {\n    9ayed i = 0;\n    ma7dBa9i(i < n) {\n        i = i + 1;\n        9ayed result = "";\n        ilakan (i % 3 == 0) {\n            result = result + "Fizz";\n        }\n        ilakan (i % 5 == 0) {\n            result = result + "Buzz";\n        }\n        ilamakanch (result == "") {\n            golih(result);\n        } sinn {\n            golih(i);\n        }\n    }\n}\n',
    },
    {
      name: 'Numerical Integration',
      code: 'golih("Calculating the integral of a function using different methods:");\ngolih("f(x) = x^2");\ngolih("a = 0, b = 1");\ngolih("n = 1000");\ngolih("");\ngolih("Rectangles method: " + methodRectangles(0, 1, 1000));\ngolih("Trapezes method: " + methodTrapezes(0, 1, 1000));\ngolih("Simpson method: " + methodSimpson(0, 1, 1000));\n\nfonksyon f(x) {\n  rjje3 (x*x);\n}\n\nfonksyon methodRectangles(a, b, n) {\n  9ayed h = (b-a)/n;\n  9ayed sum = 0;\n  9ayed i = 0;\n  ma7dBa9i(i < n) {\n    9ayed xi = a + (i*h);\n    sum = sum + (f(xi));\n    i = i + 1;\n  }\n  rjje3 (sum*h);\n}\n\nfonksyon methodTrapezes(a, b, n) {\n  9ayed h = (b-a)/n;\n  9ayed sum = (f(a) + f(b))/2;\n  9ayed i = 1;\n  ma7dBa9i(i < n) {\n    9ayed xi = a + (i*h);\n    sum = sum + f(xi);\n    i = i + 1;\n  }\n  rjje3 (sum*h);\n}\n\nfonksyon methodSimpson(a, b, n) {\n  9ayed h = (b-a)/n;\n  9ayed sum = f(a) + f(b);\n  9ayed i = 1;\n  ma7dBa9i(i < n) {\n    9ayed xi = a + (i*h);\n    sum = sum + 2*f(xi);\n    i = i + 1;\n  }\n  i = 0;\n  ma7dBa9i(i < n) {\n    9ayed xi = a + ((i+1/2)*h);\n    sum = sum + 4*f(xi);\n    i = i + 1;\n  }\n  rjje3 (sum*h/6);\n}\n',
    },
    {
      name: 'Non Linear Equation',
      code: 'fonksyon f(x) { rjje3 x*x-7; }\nfonksyon df(x) { rjje3 2*x; }\nfonksyon g(x) { rjje3 (3*x+7/x)/4; }\n\n9ayed acceptableErr = 0.0001;\n\ngolih("Calculating the root of f(x) = x^3 - 2x - 5 using different methods:");\ngolih("1. Dicotomic method: " + dicothomy(2, 3, acceptableErr));\ngolih("2. Newton\'s method: " + newton(2, acceptableErr));\ngolih("3. Point fixe method: " + pointFixe(2, acceptableErr));\n\nfonksyon dicothomy(a,b,e) {\n  ma7dBa9i (((b-a)/2) > e) {\n    9ayed mid = (a+b)/2;\n    ilakan (f(mid) == 0) {\n      rjje3 mid;\n    } sinn {\n      ilakan (f(a) * f(mid) < 0) {\n        b = mid;\n      } sinn {\n        a = mid;\n      }\n    }\n  }\n  rjje3 (a+b)/2;\n}\n\nfonksyon newton(x0, e) {\n  9ayed x1 = x0 - (f(x0)/df(x0));\n  ma7dBa9i (abs(x1-x0) > e) {\n    x0 = x1;\n    x1 = x0 - (f(x0)/df(x0));\n  }\n  rjje3 x1;\n}\n\nfonksyon pointFixe(x0, e) {\n  9ayed x1 = g(x0);\n  ma7dBa9i (abs(x1-x0) > e) {\n    x0 = x1;\n    x1 = g(x0);\n  }\n  rjje3 x1;\n}\n',
    },
  ];
}

// const rootPath = path.join(__dirname, '..', '..', '..');

// const examples = [
//   {
//     name: 'Hello, World!',
//     codePath: path.join(rootPath, 'examples/hello_world.ds'),
//   },
//   {
//     name: 'FizzBuzz',
//     codePath: path.join(rootPath, 'examples/fizz_buzz.ds'),
//   },
//   {
//     name: 'Numerical Integration',
//     codePath: path.join(rootPath, 'examples/integral.ds'),
//   },
//   {
//     name: 'Non Linear Equation',
//     codePath: path.join(rootPath, 'examples/non_linear_eqq.ds'),
//   },
// ];

// export async function getExamples(): Promise<ICodingExample[]> {
//   return await Promise.all(
//     examples.map(async (example) => {
//       const codeBuffer = await readFile(example.codePath);
//       return {
//         name: example.name,
//         code: codeBuffer.toString(),
//       };
//     })
//   );
// }
