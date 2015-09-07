import Ember from 'ember';

export default Ember.Component.extend({
    actions: {
        completeTask: function(task) {
            this.sendAction('completeTask', task);
            this.set('task', {});
        }
    }
});
