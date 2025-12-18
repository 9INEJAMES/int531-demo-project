const API_BASE = process.env.NEXT_PUBLIC_API_BASE || 'http://localhost:8080'

export async function fetchUsersClient() {
    const url = `${API_BASE}/api/users`
    const res = await fetch(url)

    if (!res.ok) {
        const text = await res.text()
        throw new Error(`fetchUsersClient failed: ${res.status} ${text}`)
    }

    return res.json()
}

export async function createUserClient(id, name) {
    const url = `${API_BASE}/api/users`
    const res = await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ id, name }),
    })

    if (!res.ok) {
        const text = await res.text()
        throw new Error(`createUserClient failed: ${res.status} ${text}`)
    }

    return res.json() // { id }
}

export async function fetchUserByIdClient(id) {
    const url = `${API_BASE}/api/users/${id}`
    const res = await fetch(url)

    if (!res.ok) {
        const text = await res.text()
        throw new Error(`fetchUserByIdClient failed: ${res.status} ${text}`)
    }

    return res.json()
}

export async function updateUserClient(id, name) {
    const url = `${API_BASE}/api/users/${id}`
    const res = await fetch(url, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name }),
    })

    if (!res.ok) {
        const text = await res.text()
        throw new Error(`updateUserClient failed: ${res.status} ${text}`)
    }
}

export async function deleteUserClient(id) {
    const url = `${API_BASE}/api/users/${id}`
    const res = await fetch(url, {
        method: 'DELETE',
    })

    if (!res.ok) {
        const text = await res.text()
        throw new Error(`deleteUserClient failed: ${res.status} ${text}`)
    }
}
