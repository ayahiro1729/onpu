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
import { useState } from "react";

export const UserSearch = () => {
  const placeholders = [
    "",
  ];

  const [query, setQuery] = useState('');
  const [suggestions, setSuggestions] = useState([
    'React', 'JavaScript', 'Typescript', 'Node.js', 'Python', 'Java', 'C#', 'Go', 'Raaa', 'aaa'
  ]);

  const filteredSuggestions = query
    ? suggestions.filter((suggestion) =>
        suggestion.toLowerCase().includes(query.toLowerCase())
      ).slice(0, 5) 
    : [];

  const handleChange = (text: string) => {
    console.log(text);
  };
  const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("submitted");
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
              onChange={(event) => setQuery(event.target.value)}
              onSubmit={onSubmit}
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
      </DialogContent>
    </Dialog>
  );
};

