import { PlaceholdersAndVanishInput } from "~/components/ui/placeholders-and-vanish-input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "~/components/ui/dialog"
import searchLogo from '/search.svg'
import { useState, ChangeEvent } from "react";

export const UserSearch = () => {
  type UserSearchResult = {
    user_id: number;
    user_name: string;
    display_name: string;
    icon_image: string;
  }

  const placeholders = [
    "",
  ];

  const [query, setQuery] = useState('');
  const [suggestions, setSuggestions] = useState([
    'React', 'JavaScript', 'Typescript', 'Node.js', 'Python', 'Java', 'C#', 'Go', 'Raaa', 'aaa'
  ]);
  const [results, setResults] = useState<UserSearchResult[]>([]);

  const filteredSuggestions = query
    ? suggestions.filter((suggestion) =>
        suggestion.toLowerCase().includes(query.toLowerCase())
      ).slice(0, 5) 
    : [];

  const handleChange = async (event: ChangeEvent<HTMLInputElement>) => {
    const searchInput: string|null = event.target.value;

    if (searchInput) {
      const searchString = encodeURIComponent(searchInput);
      const userSearchResponse = await fetch(`http://localhost:8080/api/v1/user?search_string=${searchString}`);
      if (!userSearchResponse.ok) {
        throw new Error (`Failed to fetch user search data: ${userSearchResponse.statusText}`);
      }
      const userSearchData = await userSearchResponse.json();
      setResults(userSearchData.users);
    } else {
      setResults([]);
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <img src={searchLogo} className='w-7 h-auto'/>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Search for a user</DialogTitle>
          <DialogDescription>
          </DialogDescription>
          <div className="flex flex-col justify-center items-center">
            <PlaceholdersAndVanishInput
              placeholders={placeholders}
              onChange={handleChange}
            />
          </div>
          {filteredSuggestions.length > 0 && (
            <div className="w-full bg-white border border-gray-300 rounded-md shadow-lg">
              <ul className="py-1">
                {filteredSuggestions.map((suggestion) => (
                  <li
                    key={suggestion}
                    className="px-4 py-2 text-gray-700 hover:bg-gray-100 cursor-pointer flex justify-start items-center"
                  >
                    {suggestion}
                  </li>
                ))}
              </ul>
            </div>
        )}
        </DialogHeader>
        <div>
          <ul>
            {results.length > 0 ? (
              results.map((user: UserSearchResult) => (
              <li key={user.user_id}>
                <img src={user.icon_image} className="w-10 h-10 rounded-full" />
                <span>{user.display_name}</span>
                <span>{user.user_name}</span>
              </li>
              ))) : (
                <span>no results</span>
              )}
          </ul>
        </div>
      </DialogContent>
    </Dialog>
  );
};

