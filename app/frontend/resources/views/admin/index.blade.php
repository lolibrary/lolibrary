@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row mb-3">
        <div class="col text-center">
            <h1 class="h3">Admin Dashboard</h1>
        </div>
    </div>
    <div class="card">
        <div class="card-body">
            <p>
                Welcome to the Lolibrary Admin panel —
                this is a quick stopgap since I ripped it
                from the in-development version of Lolibrary.
            </p>

            <p>
                You can edit, add, publish and preview items from
                the navigation above. It should work somewhat on mobile,
                but please try and stay on desktop for now.
            </p>

            <p>
                The current permissions are:
            </p>
            <ul>
                <li><strong>Junior Lolibrarian</strong> — You can submit draft items for review by lolibrarians and over.</li>
                <li><strong>Lolibrarian</strong> — You can submit items for without requiring review.</li>
                <li><strong>Senior Lolibrarian</strong> — You can review draft submissions and approve them, and edit existing items.</li>
                <li><strong>Moderator (admin)</strong> — You can edit users and have full access to item editing. You can also delete items.</li>
            </ul>
            <p>Your level is in the navigation bar above.</p>
        </div>
    </div>
</div>
@endsection