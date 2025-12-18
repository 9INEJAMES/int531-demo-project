import { useEffect, useState } from 'react'
import Link from 'next/link'
import { fetchUsersClient, deleteUserClient } from '../lib/api'

export default function Home() {
    const [users, setUsers] = useState([])
    const [error, setError] = useState(null)
    const [loading, setLoading] = useState(true)

    const loadUsers = async () => {
        setLoading(true)
        try {
            const data = await fetchUsersClient()
            setUsers(data || [])
            setError(null)
        } catch (e) {
            setError(e.message)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        loadUsers()
    }, [])

    const handleDelete = async (id) => {
        if (!confirm('Are you sure you want to delete this user?')) return
        try {
            await deleteUserClient(id)
            await loadUsers()
        } catch (e) {
            alert(e.message)
        }
    }

    return (
        <div className="min-h-screen py-10 px-4 sm:px-6 lg:px-8">
            <div className="max-w-4xl mx-auto">
                <div className="flex justify-between items-center mb-8">
                    <h1 className="text-3xl font-bold text-gray-900 tracking-tight">User Management</h1>
                    <Link 
                        href="/users/create"
                        className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 transition-colors"
                    >
                        Create User
                    </Link>
                </div>

                {error && (
                    <div className="rounded-md bg-red-50 p-4 mb-6">
                        <div className="flex">
                            <div className="flex-shrink-0">
                                <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                                    <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                                </svg>
                            </div>
                            <div className="ml-3">
                                <h3 className="text-sm font-medium text-red-800">Error loading users</h3>
                                <div className="mt-2 text-sm text-red-700">
                                    <p>{error}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                )}

                <div className="bg-white shadow overflow-hidden sm:rounded-lg border border-gray-200">
                    {loading ? (
                        <div className="p-12 text-center text-gray-500">Loading...</div>
                    ) : users.length === 0 ? (
                        <div className="p-12 text-center text-gray-500">No users found. Create one to get started.</div>
                    ) : (
                        <ul className="divide-y divide-gray-200">
                            {users.map((user) => (
                                <li key={user.id} className="px-6 py-4 flex flex-col sm:flex-row sm:items-center sm:justify-between hover:bg-gray-50 transition-colors duration-150 ease-in-out">
                                    <div className="flex-1 min-w-0">
                                        <p className="text-sm font-medium text-indigo-600 truncate">
                                            ID: {user.id}
                                        </p>
                                        <p className="text-lg text-gray-900 font-semibold truncate">
                                            {user.name}
                                        </p>
                                    </div>
                                    <div className="mt-4 sm:mt-0 flex items-center space-x-4">
                                        <Link 
                                            href={`/users/${user.id}`}
                                            className="inline-flex items-center px-3 py-1.5 border border-indigo-200 text-xs font-medium rounded text-indigo-700 bg-indigo-50 hover:bg-indigo-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                                        >
                                            Edit
                                        </Link>
                                        <button
                                            onClick={() => handleDelete(user.id)}
                                            className="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                                        >
                                            Delete
                                        </button>
                                    </div>
                                </li>
                            ))}
                        </ul>
                    )}
                </div>
            </div>
        </div>
    )
}
