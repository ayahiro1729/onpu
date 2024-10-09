import { PlaceholdersAndVanishInput } from "~/components/ui/placeholders-and-vanish-input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
  DialogClose,
} from "~/components/ui/dialog"
import searchLogo from '/search.svg'
import { useState, ChangeEvent } from "react";
import { Link } from "@remix-run/react";
import { UserSearchResult } from "~/types/types";

export const UserSearch = () => {
  const placeholders = [
    "",
  ];

  const [results, setResults] = useState<UserSearchResult[]>([]);

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

  const handleOpen = () => {
    setResults([]);
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <img src={searchLogo} className='w-7 h-auto' onClick={handleOpen}/>
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
              onSubmit={() => {}}
            />
          </div>
        </DialogHeader>
        <div>
          <ul className="flex flex-col gap-2">
            {results.length > 0 ? (
              results.map((user: UserSearchResult) => (
              <li key={user.user_id}>
                <Link to={`/user/${user.user_id}`}>
                  <DialogClose className="flex gap-3 items-center">
                    <img src={user.icon_image} className="w-10 h-10 rounded-full"/>
                    <span>{user.display_name}</span>
                  </DialogClose>
                </Link>
              </li>
              ))) : (
                <span className="flex justify-center text-2xl">No Results</span>
              )}
          </ul>
        </div>
      </DialogContent>
    </Dialog>
  );
};

