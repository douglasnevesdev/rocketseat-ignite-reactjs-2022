type User = {
  permissions: string[];
  roles: string[];
}

type ValidateUserPermissionsParams = {
  user: User;
  permissions?: string[];
  roles?: string[];
}

export function validateUserPermissions({ user, permissions, roles }: ValidateUserPermissionsParams) {

  //Valido se usuario tem permissao e roles, se não tiver qualquer um dos 2 declarados retorna false.
  if (permissions?.length > 0) {


    //Percorre sobre as permissões necessarias do componente e compara se o usuario tem todas as necessarias.
    const hasAllPermissions = permissions.every((permission) => {
      return user?.permissions.includes(permission);
    });


    //Se retorna true, percorre sobre todas as roles do componente e verifica se o usuario atende todas as necessidades.
    if (hasAllPermissions) {
      if (roles?.length > 0) {

        const hasAllRoles = roles.some((role) => {
          return user?.roles.includes(role);
        });

        return hasAllRoles;

      }
    }

  }

  return false;

}