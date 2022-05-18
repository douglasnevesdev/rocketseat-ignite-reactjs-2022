import { ProductItem } from "./ProductItem";
import { List, AutoSizer, ListRowRenderer } from 'react-virtualized';

type SearchResultsProps = {
  totalPrice: number;
  results: Array<{
    id: number;
    price: number;
    title: string;
    priceFormatted: string;
  }>
  onAddToWishlist: (id: number) => void;
}

export function SearchResults({ results, totalPrice, onAddToWishlist }: SearchResultsProps) {

  const rowRenderer: ListRowRenderer = ({ index, key, style }) => {
    return (
      <div key={key} style={style}>
        <ProductItem
          product={results[index]}
          onAddToWishlist={onAddToWishlist} />
      </div>
    )
  }


  return (
    <div>
      <h2>{totalPrice}</h2>

      <List
        height={300}
        rowHeight={30}
        width={900}
        overscanRowCount={5}
        rowCount={results.length}
        rowRenderer={rowRenderer}
      />

      {/*results.map(product => {
        return (
          <ProductItem
            product={product}
            key={product.id}
            onAddToWishlist={onAddToWishlist} />
        );
      })*/}
    </div>
  );
}