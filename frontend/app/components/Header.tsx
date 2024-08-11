import React from 'react'
import logo from '/OnpuLogo.jpg'
import { UserSearch } from './UserSearch'
import { Link } from '@remix-run/react'
import { HeaderProps } from '~/types/types'

export const Header :React.FC<HeaderProps> = ({ iconImage, myUserId }) => {
  return (
    <div className="p-2 flex justify-between items-center fixed top-0 left-0 right-0 bg-white border-b-2 border-current z-[9999]">
      <Link to={`/user/${myUserId}`}>
        <div className='flex flex-column gap-1'>
          <img src={logo} className="w-12 shadow rounded-full max-w-full h-auto align-middle border-none" />
          <p className='flex justify-center items-center'>Onpu</p>
        </div>
      </Link>
      <div className='flex flex-column gap-3'>
        <UserSearch />
        <Link to={`/user/${myUserId}/edit`}>
          <img src={iconImage} className='w-8 h-8 rounded-full'/>
        </Link>
      </div>
    </div>
  )
}