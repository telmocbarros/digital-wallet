import SearchIcon from '../atoms/icons/SearchIcon';
import Input from '../atoms/Input';
import './SearchBar.css';

interface SearchBarProps {
  placeholder?: string;
  onSearch?: (value: string) => void;
  className?: string;
}

export default function SearchBar({
  placeholder = 'Search',
  onSearch,
  className = '',
}: SearchBarProps) {
  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    if (onSearch) {
      onSearch(e.target.value);
    }
  }

  return (
    <div className={`search-container ${className}`}>
      <SearchIcon className="search-icon" />
      <Input
        type="text"
        placeholder={placeholder}
        onChange={handleChange}
        className="search-input"
      />
    </div>
  );
}
