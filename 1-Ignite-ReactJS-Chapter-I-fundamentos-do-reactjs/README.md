# 1. Configurando ambiente

## Introdução do módulo

- React é uma biblioteca para criação de interface.
- SPA - Single-page application.

## Criando estrutura do projeto

```jsx
yarn init -y
yarn add react
yarn add react-dom
```

- Foi criado a pasta public
- Foi criado a pasta src

## Configurando Babel

- Babel converte o código de uma maneira que todos os navegadores e ambiente possa entender nosso código.

```jsx
yarn add 
@babel/core // Biblioteca do babel
@babel/cli  // Libera as linhas de comando
@babel/preset-env //Converte o codigo da melhor maneira possivel
-D
```

- Criar arquivo babel.config.js

```jsx
module.exports = {
  presets: [
    '@babel/preset-env',
    '@babel/preset-react'
  ]
}
```

- No código abaixo transformamos o javascript usando babel, o resultado será salvo na pasta dist.

```jsx
yarn babel src/index.js --out-file dist/bundle.js
```

O Babel não sabe interpretar o react ( Javascript + HTML ), para isso instalamos outro preset

```jsx
yarn add @babel/preset-react -D
```

No arquivo de configuração do babel adicionamos o preset-react.

- Renomear os arquivos para a extensão .jsx que é a nomenclatura que utilizamos quando usamos html dentro do javascript, isso foi algo que foi introduzido pelo React, você pode tanto usar .js ou .jsx porque não existe nada que diz se isso esta errado ou não.

Nessa etapa entendemos o papel do babel, na medida em que o projeto avança, instalamos novos presets fornecidos pelo babel para interpretar as tecnologias utilizada no código como react, todas essa foram bibliotecas instalar + adicionadas no arquivo de config do babel.

## Configurando Webpack

Webpack muitas vezes é utilizado em conjunto com babel.

- Fornece uma serie de configurações que chamamos de loaders que irão ensinar nossa aplicação a maneira de tratar cada tipo de importação, exemplo: .css, .json, .jpg, .sass entre outros.
- Por fim o webpack pega todos os arquivos e converte para arquivos que são compreendidos pelo browser.

```jsx
yarn add webpack webpack-cli -D
yarn add babel-loader -D
```

- Assim como babel, o webpack precisa de seu arquivo de configuração.

```jsx
const path = require('path');

module.exports = {
  entry: path.resolve(__dirname, 'src', 'index.jsx'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['.js', '.jsx'],
  },
  module: {
    rules: [
      {
        test: /\.jsx$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      }
    ],
  }
}
```

Para executar a aplicação usando o webpack junto com babel utilize o comando:

```jsx
yarn webpack
```

Nessa etapa entendemos a utilização do webpack no projeto que normalmente é utilizado junto com o babel para que tudo possa ser interpretado (.js, .jpg, .sass entre outros) e retorna um único arquivo javascript. O webpack usa loaders que ensina a aplicação a tratar cada tipo de importação e por fim converte tudo de uma maneira que o browser entenda.

## Estrutura do ReactJS

- Cria toda interface da aplicação através do javascript.
- Quando trabalhamos com ReactJS temos um único elemento no arquivo .html (index.html) que é o root.
- Modificamos o arquivo de configuração webpack para que ele executa as paginas em React sem a necessidade de importa React em todas as paginas.

```jsx
module.exports = {
  presets: [
    '@babel/preset-env',
    ['@babel/preset-react', {
      runtime: 'automatic'
    }]
  ]
}
```

- Webpack fornece a possibilidade de você executar no modo desenvolvedor ou produção, nativamente ele é produção, ou seja, demora um pouco mais para executar porem otimiza o código, para definir qual o tipo você pode adicionar uma config na configuração do webpack cuja parâmetro se chama mode: 'product'.

## Servindo HTML estático

Existe um plugin para o webpack que injeta o arquivo javascript em nosso HTML para que não seja necessário ficar importando.

```jsx
yarn add html-webpack-plugin -d
```

Dentro das configurações do webpack adicionamos.

```jsx
plugins: [
    new HtmlWebPackPlugin({
      template: path.resolve(__dirname, 'public', 'index.html')
    })
  ]
```

## Webpack Dev Server

Existe um plugin que fica observando alterações na pasta src e cria automaticamente o build.

```jsx
yarn add webpack-dev-server -d
```

Adicionamos nas configurações do webpack.

```jsx
devServer: {
    static: {
      directory: path.resolve(__dirname, 'public'),
    }
  },
```

Agora para executar use o comando:

```jsx
yarn webpack serve
```

É necessário definir o webpack mode: 'development'.

## Utilizando source maps

Uma funcionalidade do webpack, permite visualizar o código de forma original mesmo quando o código da aplicação esta embaralhado. 

Imagina que temos um erro no código, no console do google ele vai exibir o erro do arquivo bundle, isso complica a deputação, então usamos o source maps para facilitar.

Adicionamos ao arquivo de configuração do webpack o seguinte comando:

```jsx
devtool: 'eval-source-map'
```

Esse comando acima é usado em ambiente de desenvolvimento, agora se ocorrer algum erro será mostrado a exata linha no console do google.

## Ambiente dev e produção

No arquivo de configuração do webpack adicionamos uma variável global para informar se o projeto esta em desenvolvimento ou produção.

```jsx
const path = require('path');
const HtmlWebPackPlugin = require('html-webpack-plugin');

const isDevelopment = process.env.NODE_ENV !== 'production';

module.exports = {
  mode: isDevelopment ? 'development' : 'production',
  devtool: isDevelopment ? 'eval-source-map' : 'source-map',
  entry: path.resolve(__dirname, 'src', 'index.jsx'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['.js', '.jsx'],
  },
  devServer: {
    static: {
      directory: path.resolve(__dirname, 'public'),
    }
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: path.resolve(__dirname, 'public', 'index.html')
    })
  ],
  module: {
    rules: [
      {
        test: /\.jsx$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      }
    ],
  }
}
```

Para criar as variáveis de ambiente no ReactJS precisamos executar um comando, como esse comando seria diferente no Windows/Linux/Mac usamos uma biblioteca:

```jsx
yarn add cross-env -D
```

Agora no package.json, criamos um parâmetro scripts

```jsx
"scripts": {
    "dev": "webpack serve",
    "build": "cross-env NODE_END=production webpack"
  },
```

## Importando arquivo CSS

- Para o webpack entender o .css ele precisa de loaders específicos, para isso vamos instalar.

```jsx
yarn add style-loader -d
yarn add css-loader -d
```

Adicionamos no arquivo de configuração do webpack os comandos para ele interpretar o .css

```jsx
rules: [
      {
        test: /\.jsx$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      },
      {
        test: /\.css$/,
        exclude: /node_modules/,
        use: ['style-loader', 'css-loader'],
      }
    ],
```

## Utilizando SASS

Pre-processador do css.

```jsx
yarn add sass-loader -d
yarn add node-sass -d
```

Agora no arquivo de configuração do webpack vamos modificar a regra de .css

- .sass → Não precisa tratar as {} do css.
- .scss → Precisa tratar as chaves de cada propriedade css.

```jsx
rules: [
      {
        test: /\.jsx$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      },
      {
        test: /\.scss$/,
        exclude: /node_modules/,
        use: ['style-loader', 'css-loader', 'sass-loader'],
      }
    ]
```

Uma das vantagens do Sass é o encadeamento.

# 2. Conceitos Importantes

## Primeiro componente React

- É uma função que começa com a primeira letra maiúscula e devolve um HTML.
- Por convenção temos apenas um componente por arquivo (É possível adicionar mais de um)

## Propriedades no React

### Propriedas

- Funciona como atributos funciona em TAGS HTML.
- São informações através de variáveis que podemos passar para o componente funcionar diferente.
- O Conceito seria poder enviar uma informação do componente PAI para o componente FILHO.
- Recebemos as propriedades no componente FILHO por um parâmetro chamado props.

{  props.repository  ??  'Default'  } // Se não tiver nada em uma propriedade ele usa o texto Defaut

- Para evitar passar vários parâmetros para um componente podemos criar um objeto e depois passar como parâmetro desse componente.

```jsx
const repository = {
  name: 'unform',
  description: 'Form in React',
  link: 'https://github.com/unform/unform'
}

<RepositoryItem repository={repository} />
```

E dentro do componente:

```jsx
export function RepositoryItem(props) {
  return (
    <li>
      <strong>{props.repository?.name ?? 'Default'}</strong>
      <p>{props.repository?.description}</p>
      <a href={props.repository?.link}>
        Acessar Repositório
      </a>
    </li>
  );
}
```

## Estado do componente

- São variáveis que o React vai monitorar mudança.
- Para utilizar utilize import useState(); da biblioteca react

```jsx
const [counter, setCounter] = useState(0); // Exemplo
```

- Podemos usar variáveis comuns, porem não será refletido na interface.

## A imutabilidade no React

Imutabilidade é um conceito da computação.

- Não podemos alterar diretamente o conteúdo da variável, ou seja, reconstruímos uma nova variável com o valor novo, por isso usamos o SetCounter(counter + 1);

```jsx
novoUsuario = [...usuarios, 'Douglas'];
```

## Fast Refresh no WebPack

- Toda vez que alteramos o código, o webpack reseta o código para o total zero.
- O React criou uma biblioteca (plugin) chamada react-refresh-webpack-plugin aonde não perdemos as informações, exemplo, um carrinho de compras esta cheio de itens mais você precisa alterar um codigo, mesmo o webpack dando reload a aplicação vai manter os dados do carrinho.

```jsx
yarn add @pmmmwh/react-refresh-webpack-plugin@0.5.0-rc.4
```

Quando temos apenas o IF, podemos usar &&

# 3. Chamadas HTTP

## Utilizando o useEffect

- Dispara uma função quando algo acontecer na aplicação

```jsx
useEffect(()=>{}, []); // Segundo parametro é as dependencias, ou seja, quais informações que se mudarem o useEffect deve executar novamente.
```

- Não esquecer de adicionar o segundo parâmetro [], se esquecer de declarar ele ficara em loop.

## Listando Repositórios

Exemplo de fetch

```jsx
useEffect(() => {
    fetch('https://api.github.com/users/douglasneves-net/repos')
      .then(response => response.json())
      .then(data => setRepositories(data));
  }, []);
```

# 4. Usando TypeScript

## Fundamentos do TypeScript

SuperSet → Conjunto de funcionalidades que adicionamos encima de uma linguagem.

- Podemos usar Type ou Interface, de preferencia use Type para tipos e interface quando precisamos de um contrato.

```jsx
Type User = {
	name: string;
	email: string;
}
```

## TypeScript no ReactJS

Para adicionar o TypeScript ao projeto e criar o arquivo de configuração.

```jsx
yarn add typescript -D // Adicionar TypeScript no projeto
yarn add tsc --init // Criar arquivo de configuração
```

- No arquivo de configuração do Typescript foram removidos alguns comentarios.

```jsx
{
  "compilerOptions": {
    "lib": [
      "dom",
      "dom.iterable",
      "esnext"
    ],
    "allowJs": true, //Informa que teremos arquivos javascript e Typescript.
    "jsx": "react-jsx", //informa que estamos usando JSX do React
    "noEmit": true, 
    "strict": true,
    "moduleResolution": "node",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "allowSyntheticDefaultImports": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true
  },
  "include": [
    "src"
  ] //Informamos que todo codigo da nossa aplicação esta em src.
}
```

- O webpack por padrão esta usando o babel-loader, porem o babel-loader não sabe interpretar código Typescript, apenas javascript, então precisamos adicionar mais uma biblioteca ao babel para ele saber interpretar Typescript.

```jsx
yarn add @babel/preset-typescript -D
```

- Agora no arquivo de configuração do babel adicionamos o preset.

```jsx
module.exports = {
  presets: [
    '@babel/preset-env',
    '@babel/preset-typescript',
    ['@babel/preset-react', {
      runtime: 'automatic'
    }]
  ]
}
```

- Agora no webpack ajustaremos para ler tanto .jsx como .tsx.

```jsx
const path = require('path');
const HtmlWebPackPlugin = require('html-webpack-plugin');
const ReactRefreshWebPackPlugin = require('@pmmmwh/react-refresh-webpack-plugin');

const isDevelopment = process.env.NODE_ENV !== 'production';

module.exports = {
  mode: isDevelopment ? 'development' : 'production',
  devtool: isDevelopment ? 'eval-source-map' : 'source-map',
  entry: path.resolve(__dirname, 'src', 'index.tsx'),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['.js', '.jsx', '.ts', '.tsx'],
  },
  devServer: {
    static: {
      directory: path.resolve(__dirname, 'public'),
    },
    hot: true,
  },
  plugins: [
    isDevelopment && new ReactRefreshWebPackPlugin(),
    new HtmlWebPackPlugin({
      template: path.resolve(__dirname, 'public', 'index.html')
    })
  ].filter(Boolean),
  module: {
    rules: [
      {
        test: /\.(j|t)sx$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            plugins: [
              isDevelopment && require.resolve('react-refresh/babel')
            ].filter(Boolean)
          }
        }
      },
      {
        test: /\.scss$/,
        exclude: /node_modules/,
        use: ['style-loader', 'css-loader', 'sass-loader'],
      }
    ],
  }
}
```

- Algumas bibliotecas não inclui as definições de tipos do Typescript, ou seja, não inclui toda a parte que o Typescript precisa para entender como a biblioteca funciona. Sendo assim a definição de tipos pode ser criado pela comunidade ou esta a parte na biblioteca.

```jsx
yarn add @types/react-dom -d
```

## Componentes com TypeScript

- Existe uma funcionalidade do Javascript chamada Generic, ou seja, um tipo genérico, utilizamos <> para representar.

# 5. Finalizando aplicação

## Utilizando React DevTools

Instalar extensão React DevTools.

## Finalização do módulo