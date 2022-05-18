import { Flex, Button, Stack, FormLabel, FormControl } from '@chakra-ui/react';
import { Input } from '../components/Form/Input';
import { SubmitHandler, useForm } from 'react-hook-form';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';


type SignInFormData = {
  email: string;
  password: string;
}

const signInFormSchema = yup.object().shape({
  email: yup.string().required("E-mail obrigatório").email("E-mail inválido"),
  password: yup.string().required("Senha obrigatória"),
});


export default function SignIn() {

  //Integra nosso yup com react-hook-form e passa o schema do yup.
  const { register, handleSubmit, formState } = useForm({
    resolver: yupResolver(signInFormSchema)
  });

  const { errors } = formState;


  //Poderia colocar o tipo diretamente em values.
  //Usei a função SubmitHandler porque ela recebe como função values e o event de submit do formulario, ou seja, ele deixa o evento tipado e podemos usar para alguma ação na função.
  const handleSignIn: SubmitHandler<SignInFormData> = (values) => {
    console.log(values);
  }

  return (
    <Flex
      width='100vw'
      height='100vh'
      alignItems='center'
      justifyContent='center'
      //Se usar o handleSignIn diretamente não consigo obter os dados do form, então adicionamos o handleSubmit em volta de tudo.
      onSubmit={handleSubmit(handleSignIn)}
    >
      <Flex
        as='form'
        width="100%"
        maxWidth={360}
        bg="gray.800"
        padding='8'
        borderRadius={8}
        flexDir="column"
      >
        <Stack spacing='4'>
          <Input type="email" name="email" label="E-mail" {...register('email')} error={errors.email} />
          <Input type="password" name="password" label="Senha" {...register('password')} error={errors.password} />
        </Stack>

        <Button type="submit" marginTop="6" colorScheme="red" size="lg" >Entrar</Button>
      </Flex>
    </Flex>
  )
}
