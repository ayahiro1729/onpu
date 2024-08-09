import React from 'react'
import logo from '/OnpuLogo.jpg'
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar"
import { UserSearch } from './UserSearch'
import { Link } from '@remix-run/react'

export const Header = () => {
  return (
    <div className="p-2 flex justify-between items-center fixed top-0 left-0 right-0 bg-white border-b-2 border-current z-[9999]">
      <div className='flex flex-column gap-1'>
        <img src={logo} className="w-12 shadow rounded-full max-w-full h-auto align-middle border-none" />
        <p className='flex justify-center items-center'>Onpu</p>
      </div>
      <div className='flex flex-column gap-3'>
        <UserSearch />
        <Link to="/profileedit">
          <Avatar className='w-8 h-auto'>
            <AvatarImage src="https://github.com/shadcn.png" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
        </Link>
      </div>
    </div>
  )
}