#!/bin/bash

# =============================================================================
#
# Darklist
# https://github.com/h0lysp4nk/Darklist
#
# =============================================================================

echo '    ____             __   ___      __  '
echo '   / __ \____ ______/ /__/ (_)____/ /_ '
echo '  / / / / __ `/ ___/ //_/ / / ___/ __/ '
echo ' / /_/ / /_/ / /  / ,< / / (__  ) /_   '
echo '/_____/\__,_/_/  /_/|_/_/_/____/\__/   '
echo '                                       '
echo ''
echo '==================================================='
echo
echo

# Set dir
script_dir=`dirname $0`
cd $script_dir/dldata

# Handle cli input
case $1 in
	"ps")
		# Get the status of Darklist
		docker-compose ps

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Successfully received the status of Darklist!"
			exit 0
		else
			# Failure
			echo "[!] Couldn't get the status of Darklist!"
			exit 1
		fi;;
	"restart")
		# Restart Darklist
		docker-compose stop && docker-compose up -d

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Darklist restarted successfully!"
			exit 0
		else
			# Failure
			echo "[!] Darklist failed to restart!"
			exit 1
		fi;;
	"stop")
		# Stop Darklist
		docker-compose stop

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Darklist stopped successfully!"
			exit 0
		else
			# Failure
			echo "[!] Darklist failed to stop!"
			exit 1
		fi;;
	"start")
		# Start Darklist
		docker-compose up -d

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Darklist started successfully!"
			exit 0
		else
			# Failure
			echo "[!] Darklist failed to start!"
			exit 1
		fi;;
	"uninstall")
		# Uninstall Darklist
		docker-compose down

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			echo "[!] Darklist containers were successfully destroyed!"
			echo "[!] To finish the removal of Darklist type: 'docker system prune' ALTHOUGH be careful as this will remove any dangling containers/images! Ensure that before running this command that all you're docker containres are online!"
		else
			echo "[!] Darklist containers couldn't be destroyed!"
			exit 1
		fi;;
	"install")
		# Install Darklist
        docker network create darklist

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Darklist docker network created!"
		else
			# Failure
			echo "[!] Darklist docker network couldn't be created! Does it already exist?"
		fi

		docker-compose up -d --build

		# Check to see if that command executed successfully
		if [ $? -eq 0 ]; then
			# Success
			echo "[!] Darklist was successfully started!"
		else
			# Failure
			echo "[!] Darklist couldn't be installed!"
			exit 1
		fi;;
	*)
		echo "[!] Usage: ./darklist.sh <start|stop|restart|ps|install>";;
esac

