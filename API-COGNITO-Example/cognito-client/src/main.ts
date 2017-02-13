/// <reference path='../node_modules/amazon-cognito-identity-js/index.d.ts' />
/// <reference path='../typings/index.d.ts' />

import { CognitoUserPool, CognitoUserAttribute, CognitoUser, AuthenticationDetails, CognitoUserSession } from 'amazon-cognito-identity-js';

// YOU HAVE TO FILL THESE CONSTANTS OR PASS THEM IN ENV
let userPoolId = ''
let clientId   = ''

if (userPoolId.length === 0) {
	userPoolId = process.env.USER_POOL_ID
}
if (clientId.length === 0) {
	clientId = process.env.CLIENT_ID
}

function getUserPool() : CognitoUserPool {
	return new CognitoUserPool({
		UserPoolId: userPoolId,
		ClientId  : clientId,
	})
}

function getUsername(email: string) : string {
	return email.replace('@', '-at-').replace('.', '-dot-')
}

function getUser(userPool: CognitoUserPool, userName?: string) : CognitoUser {
	return new CognitoUser({
		Pool    : userPool,
		Username: userName,
	})
}

function getAuthenticationDetails(userName: string, password: string) : AuthenticationDetails {
	return new AuthenticationDetails({
		Username: userName,
		Password: password,
	})
}

function callback(err, result) {
	if (err) {
		log(err)
		return
	}
	log(result)
}

function signUp(email: string, password: string) {
	const userPool = getUserPool()
	const attrs: CognitoUserAttribute[] = [
		new CognitoUserAttribute({
			Name: 'email',
			Value: email,
		}),
	]
	const userName = getUsername(email)
	userPool.signUp(userName, password, attrs, null, callback)
}

function confirmRegistration(email: string, confirmationCode: string) {
	const userPool = getUserPool()
	const userName = getUsername(email)
	const user = getUser(userPool, userName)
	user.confirmRegistration(confirmationCode, true, callback)
}

function resendConfirmationCode(email: string) {
	const userPool = getUserPool()
	const userName = getUsername(email)
	const user = getUser(userPool, userName)
	user.resendConfirmationCode(callback)
}

function signIn(email: string, password: string) {
	const userPool = getUserPool()
	const userName = getUsername(email)
	const user = getUser(userPool, userName)
	const authenticationDetails = getAuthenticationDetails(userName, password)
	user.authenticateUser(authenticationDetails, {
		onSuccess: function (session: CognitoUserSession) {
			console.log('authenticateUser:success')
			log(session)
		},
		onFailure: function (err) {
			console.log('authenticateUser:failure')
			log(err)
		},
		newPasswordRequired: function (userAttributes: any, requiredAttributes: any) {
			console.log('authenticateUser:newPasswordRequired')
		},
		mfaRequired: function (challengeName: any, challengeParameters: any) {
			console.log('authenticateUser:mfaRequired')
		},
		customChallenge: function (challengeParameters: any) {
			console.log('authenticateUser:customChallenge')
		},
	})
}

const fns : Function[] = [
	usage,
	signUp,
	confirmRegistration,
	resendConfirmationCode,
	signIn,
];

function route(args: string[]) {
	if (args.length === 0) {
		usage()
		return
	}
	for (let fn of fns) {
		if (args[0] === fn.name) {
			if (args.length !== fn.length+1) {
				console.log(`Parameters mismatch: ${fn.toString().split('\n')[0]}`)
				return
			}
			fn(...args.slice(1))
			return
		}
	}
	usage()
}

function usage() {
	console.log(`Available commands are: ${fns.map((v: Function) : string => {return v.name})}`)
}

function log(value: any) {
	console.log(JSON.stringify(value, null, 4))
}

function main() {
	route(process.argv.slice(2))
}

main()
