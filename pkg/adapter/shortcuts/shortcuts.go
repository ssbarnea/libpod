package shortcuts

import (
	"github.com/containers/libpod/libpod"
	"github.com/sirupsen/logrus"
)

// GetPodsByContext returns a slice of pods. Note that all, latest and pods are
// mutually exclusive arguments.
func GetPodsByContext(all, latest bool, pods []string, runtime *libpod.Runtime) ([]*libpod.Pod, error) {
	var outpods []*libpod.Pod
	if all {
		return runtime.GetAllPods()
	}
	if latest {
		p, err := runtime.GetLatestPod()
		if err != nil {
			return nil, err
		}
		outpods = append(outpods, p)
		return outpods, nil
	}
	var err error
	for _, p := range pods {
		pod, e := runtime.LookupPod(p)
		if e != nil {
			// Log all errors here, so callers don't need to.
			logrus.Debugf("Error looking up pod %q: %v", p, e)
			if err == nil {
				err = e
			}
		} else {
			outpods = append(outpods, pod)
		}
	}
	return outpods, err
}

// GetContainersByContext gets pods whether all, latest, or a slice of names/ids
// is specified.
func GetContainersByContext(all, latest bool, names []string, runtime *libpod.Runtime) (ctrs []*libpod.Container, err error) {
	var ctr *libpod.Container
	ctrs = []*libpod.Container{}

	if all {
		ctrs, err = runtime.GetAllContainers()
	} else if latest {
		ctr, err = runtime.GetLatestContainer()
		ctrs = append(ctrs, ctr)
	} else {
		for _, n := range names {
			ctr, e := runtime.LookupContainer(n)
			if e != nil {
				// Log all errors here, so callers don't need to.
				logrus.Debugf("Error looking up container %q: %v", n, e)
				if err == nil {
					err = e
				}
			} else {
				ctrs = append(ctrs, ctr)
			}
		}
	}
	return
}
