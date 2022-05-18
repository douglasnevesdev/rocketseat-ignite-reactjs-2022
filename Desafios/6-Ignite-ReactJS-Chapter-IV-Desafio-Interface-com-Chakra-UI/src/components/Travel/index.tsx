import { Container } from "../Container";
import { Item } from "./Item";

export function Travel() {
  return (
    <Container wrap="wrap" justifyContent="space-between" mt="6rem" >
      <Item src="/travel/cocktail.png" text="vida Noturna" />
      <Item src="/travel/surf.png" text="praia" />
      <Item src="/travel/building.png" text="moderno" />
      <Item src="/travel/museum.png" text="clássico" />
      <Item src="/travel/earth.png" text="e mais" />
    </Container>
  );
}