# coding: utf-8
import os
import sysconfig
from setuptools import setup, find_packages
from setuptools.command.build_py import build_py as _build_py
from Cython.Build import cythonize


def getPaths(root_dir, exclude_files=[]):
    """get filepaths for compilation"""
    paths = []
    for root, dirs, files in os.walk(root_dir):
        for filename in files:
            if os.path.splitext(filename)[1] != '.py':
                continue

            file_path = os.path.join(root, filename)
            if file_path in exclude_files:
                continue

            paths.append(file_path)
    return paths


class build_py(_build_py):

    def find_package_modules(self, package, package_dir):
        ext_suffix = sysconfig.get_config_var('EXT_SUFFIX')
        modules = super().find_package_modules(package, package_dir)
        filtered_modules = []
        for (pkg, mod, filepath) in modules:
            if os.path.exists(filepath.replace('.py', ext_suffix)):
                continue
            filtered_modules.append((pkg, mod, filepath, ))
        return filtered_modules


setup(
    name="mibandPulse",
    version="0.1",
    scripts=['milib/mibandPulse'],
    description="A client to connect to the Miband and fetch the information and send it to pulse server",
    url="https://github.com/ramantehlan/pulse",
    packages=find_packages(),
    ext_modules=cythonize(
        getPaths('milib', ['milib/mibandPulse']),
        compiler_directives={'language_level': 3}
    ),
    cmdclass={
        'build_py': build_py
    }
)
