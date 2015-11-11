### Spanner Usecase:
## Ruby

### Build

```toml
# Assemblyfile

[application]
name = 'Example App'
repo = 'myrepo.example.com/foo/example_app'

[build]
builder = 'ruby'
version = '2.2.3'
```

* Uses the [build] section of the Assemblyfile
* Builder & Version is specified, but must be parsed earlier to choose the correct builder, we should validate it as the builder starts.

* Default Behaviour
  * Copy source from `/var/assemblyline/src/` to `/app/`
  * Check on the Presence of Gemfile and Gemfile.lock
  * Restore any cached version of `vendor/bundler`
  * Run `bundle install` with `--path vendor/bundle`
  * Run `bundle clean`
  * Save the cache

#### Proposed Config

```toml
[builder]
builder = 'ruby'
version = '2.2.3'

[[step]]
require_files = [
  'Gemfile',
  'Gemfile.lock'
]
cache = 'vendor/bundle'
script = [
  'bundle install -j4 -r3 --path vendor/bundle',
  'bundle clean'
]

[artifact]
type = 'container'
target = 'self'
```

#### Proposed Interface
* Builder config is present in `/etc/assemblyline/spanner/config.toml`
* Source to build is mounted `ro` in `/var/assemblyline/src/`
* File Cache is mounted `rw` in `/var/assemblyline/cache/`
* Execute `spanner build` inside the container