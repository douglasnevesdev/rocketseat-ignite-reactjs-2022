import { MovieProps } from "../App";
import { Header } from './Header';
import { MovieCard } from '../components/MovieCard';

interface ContentItemProps {
  selectedGenre: string;
  movies: MovieProps[];
};

export const Content = (props: ContentItemProps) => {
  // Complete aqui
  return (
    <div className="container">
      <Header title={props.selectedGenre} />

      <main>
        <div className="movies-list">
          {props.movies.map(movie => (
            <MovieCard
              title={movie.Title}
              poster={movie.Poster}
              runtime={movie.Runtime}
              rating={movie.Ratings[0].Value}
            />
          ))}
        </div>
      </main>
    </div>
  );
}