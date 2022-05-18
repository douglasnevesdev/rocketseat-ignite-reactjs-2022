import { memo, useState } from 'react';
import { AddProductToWishilistProps } from './AddProductToWishilist';
import dynamic from 'next/dynamic';

//Importando com a função dynamic, o next.js utiliza o lazy load, ou seja, so carrega quando precisamos o arquivo.
//Lazy load serve tanto para componente como para funções.
const AddProductToWishilist = dynamic<AddProductToWishilistProps>(() => {
  return import('./AddProductToWishilist').then(mod => mod.AddProductToWishilist);
}, {
  loading: () => <span>Carregando...</span>
});


interface ProductItemProps {
  product: {
    id: number;
    price: number;
    title: string;
    priceFormatted: string;
  },
  onAddToWishlist: (id: number) => void;
}

function ProductItemComponente({ product, onAddToWishlist }: ProductItemProps) {

  const [isAddingToWishlist, setIsAddingToWishlist] = useState(false);

  return (
    <div>
      {product.title} - <strong>{product.priceFormatted}</strong>
      <button onClick={() => { setIsAddingToWishlist(true) }} >Adicionar aos favoritos</button>

      {isAddingToWishlist && (
        <AddProductToWishilist
          onAddToWishlist={() => onAddToWishlist(product.id)}
          onRequestClose={() => setIsAddingToWishlist(false)}
        />
      )}

    </div>
  )
}

//Vai comparar as propriedades do objeto, neste caso product e não vai renderizar se não houver mudança
export const ProductItem = memo(ProductItemComponente, (prevProps, nextProps) => {
  //Retornando true o componente não renderiza.
  return Object.is(prevProps.product, nextProps.product)
});