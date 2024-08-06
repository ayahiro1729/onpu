import React from 'react'
import logo from '/OnpuLogo.jpg'
import searchLogo from '/search.svg'
import { Avatar, AvatarFallback, AvatarImage } from "~/components/ui/avatar"

export const Header = () => {
  return (
    <div className="p-2 flex justify-between items-center fixed top-0 left-0 right-0 bg-white border-b-2 border-current">
      <div className='flex flex-column gap-1'>
        <img src={logo} className="w-12 shadow rounded-full max-w-full h-auto align-middle border-none" />
        <p className='flex justify-center items-center'>Onpu</p>
      </div>
      <div className='flex flex-column gap-3'>
        <img src={searchLogo} className='w-7 h-auto'/>
        <Avatar className='w-8 h-auto'>
          <AvatarImage src="https://github.com/shadcn.png" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
      </div>
    </div>
  )
}