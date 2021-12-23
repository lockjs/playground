<?php

use Slim\App;

use Psr\Http\Message\ResponseInterface;
use Psr\Http\Message\ServerRequestInterface;

return function (App $app) {
    $app->get('/', \App\Action\Home::class);

    $app->post('/users', \App\Action\UserCreateAction::class);
};
