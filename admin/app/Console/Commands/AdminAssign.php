<?php

namespace App\Console\Commands;

use App\User;
use App\Role;
use Illuminate\Console\Command;

class AdminAssign extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'admin:assign {email : The user\'s email }';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Turn an existing user into an admin';

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $email = $this->argument('email');

        if (!$email) {
            $this->error('No email address');

            return;
        }

        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            $this->error('Invalid email address');

            return;
        }

        $this->assignAdmin($email);
    }

    private function assignAdmin(String $email) {
        $role = Role::where('name', 'admin')->first();

        if ($role === null) {
            $this->error('Admin role does not exist');

            return;
        }

        $user = User::where('email', $email)->first();

        if ($user === null) {
            $this->error('The user does not exist');

            return;
        }

        if ($user->role_id === $role->id) {
            $this->error('The user is already an admin');

            return;
        }

        $user->role_id = $role->id;

        if (!$user->save()) {
            $this->error('Error assigning admin role');

            return;
        }

        $this->info('Admin role assigned successfully');
    }
}
