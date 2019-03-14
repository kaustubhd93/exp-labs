import time
import logging


class PyLogger():
    """
    This can be use to log all the files in agent.log file.
    """
    def py_logger(self, logData, logType='debug', name='monitoring', rotate=True):
        if rotate:
            timestr = time.strftime("%Y%m%d%H")
            logFileName = '/var/log/' + name + '.' + timestr + ".log"
        else:
            logFileName = '/var/log' + name + '.log'
        formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
        logger = logging.getLogger(name)

        if logType == 'debug':
            level = logging.DEBUG
        else:
            level = logging.ERROR

        handler = logging.FileHandler(logFileName, mode='a')
        handler.setFormatter(formatter)

        logger.setLevel(level)
        logger.addHandler(handler)

        if logType == 'debug':
            logger.debug(logData)
        else:
            logger.exception(logData)

        logger.handlers = []
